package iteration

import "strings"

func Repeat(character string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}

func RepeatSB(character string, n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(character)
	}
	return sb.String()
}
