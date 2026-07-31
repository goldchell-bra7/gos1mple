package text

import (
	"strings"
	"unicode/utf8"
)

// stringCycle creates a string consisting of the specified number of spaces.
//
// It is used internally by Header to generate left and right padding,
// allowing the text to be centered within the requested width.
func stringCycle(width int) string {
	var builder strings.Builder

	// Build the padding one space at a time.
	// strings.Builder is used to efficiently construct the final string.
	for i := 1; i <= width; i++ {
		builder.WriteString(" ")
	}

	return builder.String()
}

// Header creates a simple text header surrounded by horizontal separators.
//
// The function centers the provided text within the specified width and
// places a line of repeated symbols above and below it, producing output
// similar to:
//
//	====================
//	     Hello World
//	====================
//
// UTF-8 rune counting is used instead of len() so that Unicode characters
// are measured correctly when calculating the text's visual width.
func Header(text, symbol string, width int) string {
	// Calculate the amount of padding required on each side of the text.
	// If the total padding is odd, the extra space is added to the right.
	left := (width - utf8.RuneCountInString(text)) / 2
	right := width - utf8.RuneCountInString(text) - left

	// Construct the header:
	//   1. Top border.
	//   2. Centered text with left and right padding.
	//   3. Bottom border.
	return strings.Repeat(symbol, width) +
		"\n" +
		stringCycle(left) +
		text +
		stringCycle(right) +
		"\n" +
		strings.Repeat(symbol, width)
}
