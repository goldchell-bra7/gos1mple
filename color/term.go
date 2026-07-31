package color

// Term wraps the provided string with the specified ANSI color escape
// sequence and automatically resets the terminal formatting afterward.
//
// The colorConst argument should be one of the color constants defined
// in this package (for example, Red, BrightBlue, or FillGreen).
// Resetting the formatting prevents the selected color from affecting
// any text printed after the returned string.
func Term(text, colorConst string) string {
	// Apply the requested ANSI color and restore the default
	// terminal formatting after the text.
	return colorConst + text + "\033[0m"
}
