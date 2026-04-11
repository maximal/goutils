package notify

import (
	"context"
	"errors"
	"net"
	"os"
	"strings"
	"time"
)

// InitSystemdNotify initializes systemd notifications.
//
// It notifies systemd on each tick of interval with the string returned by callback.
// Multiple lines in the returned string may be separated by '\n'.
//
// Returns an error if callback is nil, or if NOTIFY_SOCKET is unset (the process was not
// started by systemd with notification support).
//
// The mandatory callback is invoked on every tick.
func InitSystemdNotify(ctx context.Context, interval time.Duration, callback func() string) error {
	if callback == nil {
		return errors.New("no notification callback")
	}

	address := getValidAddress()
	if address == "" {
		return errors.New("NOTIFY_SOCKET is unset; not running under systemd with notifications")
	}

	// Tell systemd we’re ready (optional)
	if err := sendSocketNotify(address, "READY=1"); err != nil {
		return err
	}

	// If the notification socket is ready, start the ticker, and notify systemd
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				// Context done, exiting
				// ... ... ...
				return

			case <-ticker.C:
				// Ticker ticked, get message, and notify the socket
				text := strings.TrimSpace(callback())
				if text != "" {
					// systemd expects a STATUS= prefix; newlines in text are sent as-is.
					_ = sendSocketNotify(address, "STATUS="+text+"\n")
				}
			}
		}
	}()

	return nil
}

// sendSocketNotify sends a message to the systemd notification socket by its address.
func sendSocketNotify(address, message string) error {
	socket, err := net.Dial("unixgram", address)
	if err != nil {
		return err
	}
	defer func() { _ = socket.Close() }()

	_, err = socket.Write([]byte(message))
	return err
}

// getValidAddress returns systemd notification socket address.
func getValidAddress() string {
	address := os.Getenv("NOTIFY_SOCKET")
	if address == "" {
		return ""
	}

	// Abstract namespace when starts with @ (replace with \x00)
	if strings.HasPrefix(address, "@") {
		address = "\x00" + address[1:]
	}

	return address
}
