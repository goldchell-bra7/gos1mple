package text

// Bold wraps the provided string with ANSI escape sequences that enable
// bold text formatting in terminals that support ANSI colors and styles.
//
// The formatting is automatically reset at the end of the string to
// prevent subsequent terminal output from remaining bold.
func Bold(prompt string) string {
	// "\033[1m" enables bold text.
	// "\033[0m" resets all text formatting back to the terminal default.
	return "\033[1m" + prompt + "\033[0m"
}
