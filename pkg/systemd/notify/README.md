# systemd/notify
Enables systemd notifications for long-running services on Linux.

## Example

### Periodic Notifications
```go
// Notify the OS about the process state every 10 seconds.
err := notify.InitSystemdNotify(context.Background(), 10*time.Second, func() string {
	return fmt.Sprintf("Requests: %d at %.1f req/s", state.requestsTotal, state.requestsRate)
})
if err != nil {
	println("Unable to initialize systemd notifications: " + err.Error())
}
```

Then, if you run `systemctl status your-service`, you’ll get something like:
```plain
● your-service.service - Your service description
     Loaded: loaded (/usr/lib/systemd/system/your-service.service; enabled; preset: enabled)
     Active: active (running) since Sat 2026-04-11 12:08:01 UTC; 19min 26s ago
   Main PID: 777 (your-service)
     Status: "Requests: 777888 at 666.7 req/s"
     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      Tasks: 3 (limit: 28785)
     Memory: 15.6M (peak: 16.9M)
        CPU: 132ms
     CGroup: /system.slice/your-service.service
             └─777 /opt/your-service/your-service
```
Note the `Status` line.

Your `*.service` file must specify `Type=notify` to enable this functionality:
```plain
[Service]
Type=notify
ExecStart=... ... ...
... ... ...
```

### Send Once
You can send a one-time notification instead of setting up the periodic ones:
```go
err := notify.Send("Hello there")
if err != nil {
	println("Unable to send systemd notification: " + err.Error())
}
```

## Author
https://github.com/maximal
