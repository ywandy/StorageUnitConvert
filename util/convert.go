package util

import (
	"fmt"
)

const (
	_         = iota
	FLAG_Byte = iota * 10
	FLAG_Kilobyte
	FLAG_Megabyte
	FLAG_Gigabyte
	FLAG_Trillionbyte
	FLAG_Petabyte
	FLAG_Exabyte
)

func UnitConvertAuto(raw float64) string {
	suffix := ""
	b := raw
	if raw > (1 << FLAG_Exabyte) {
		suffix = "EB"
		b = raw / (1 << FLAG_Exabyte)
	} else if raw > (1 << FLAG_Petabyte) {
		suffix = "PB"
		b = raw / (1 << FLAG_Petabyte)
	} else if raw > (1 << FLAG_Trillionbyte) {
		suffix = "T"
		b = raw / (1 << FLAG_Trillionbyte)
	} else if raw > (1 << FLAG_Gigabyte) {
		suffix = "G"
		b = raw / (1 << FLAG_Gigabyte)
	} else if raw > (1 << FLAG_Megabyte) {
		suffix = "M"
		b = raw / (1 << FLAG_Megabyte)
	} else if raw > (1 << FLAG_Kilobyte) {
		suffix = "K"
		b = raw / (1 << FLAG_Kilobyte)
	} else if raw > (1 << FLAG_Byte) {
		suffix = "B"
		b = raw / (1 << FLAG_Byte)
	}
	return fmt.Sprintf("%f%s\n", b, suffix)
}
