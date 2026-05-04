package format

import "fmt"

var (
	SI_PREFIXES  = [...]string{"", "k", "M", "G", "T", "P", "E", "Z", "Y", "R", "Q"}
	IEC_PREFIXES = [...]string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi", "Yi", "Ri", "Qi"}
)

func SiUint(value uint64, unit string) string {
	if value < 1000 {
		return fmt.Sprintf("%d %s", value, unit)
	}
	var index uint8 = 0
	number := float64(value)
	for number > 998 && index < 10 {
		number /= 1000.0
		index++
	}
	return valueWithUnit(number, SI_PREFIXES[index]+unit)
}

func SiFloat(value float64, unit string) string {
	if value < 999.5 {
		return valueWithUnit(value, unit)
	}
	var index uint8 = 0
	number := value
	for number > 998 && index < 10 {
		number /= 1000.0
		index++
	}
	return valueWithUnit(number, SI_PREFIXES[index]+unit)
}

func SiBytes(bytes uint64) string {
	if bytes < 1000 {
		return fmt.Sprintf("%d B", bytes)
	}
	var index uint8 = 0
	number := float64(bytes)
	for number > 998 && index < 10 {
		number /= 1000.0
		index++
	}
	return valueWithUnit(number, SI_PREFIXES[index]+"B")
}

func IecBytes(bytes uint64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%d B", bytes)
	}
	var index uint8 = 0
	number := float64(bytes)
	for number > 998 && index < 10 {
		number /= 1024.0
		index++
	}
	return valueWithUnit(number, IEC_PREFIXES[index]+"B")
}

func valueWithUnit(value float64, unit string) string {
	separator := " "
	if unit == "" {
		separator = ""
	}
	// Three significant digits, trim trailing zeros
	if value < 1 {
		return fmt.Sprintf("%.2g%s%s", value, separator, unit)
	}
	if value < 100 {
		return fmt.Sprintf("%.3g%s%s", value, separator, unit)
	}
	// Only integer part, no decimal places
	return fmt.Sprintf("%.0f%s%s", value, separator, unit)
}
