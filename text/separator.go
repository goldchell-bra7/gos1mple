package text

import "strings"

// Separator creates a horizontal separator by repeating the specified
// symbol a given number of times.
//
// This function is useful for visually dividing sections of console output,
// such as menus, headers, or log entries.
func Separator(symbol string, width int) string {
	// Repeat the provided symbol until the requested width is reached.
	return strings.Repeat(symbol, width)
}
