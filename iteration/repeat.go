package iteration

import "strings"

func Repeat1(s string, repeatCount int) string {
	var new_s strings.Builder
	for range repeatCount {
		new_s.WriteString(s)
	}
	return new_s.String()
}

func Repeat2(s string, repeatCount int) string {
	var new_s string
	for range repeatCount {
		new_s += s
	}
	return new_s
}
