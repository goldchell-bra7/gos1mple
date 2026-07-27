package text

import (
	"strings"
	"unicode/utf8"
)

func stringCycle(width int) string {
	var builder strings.Builder

	for i := 1; i <= width; i++ {
		builder.WriteString(" ")
	}

	return builder.String()
}

func Header(text, symbol string, width int) string {
	left := (width - utf8.RuneCountInString(text)) / 2
	right := width - utf8.RuneCountInString(text) - left

	return strings.Repeat(symbol, width) + "\n" + stringCycle(left) + text + stringCycle(right) + "\n" + strings.Repeat(symbol, width)
}
