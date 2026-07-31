package text

// Cursive wraps the provided string with ANSI escape sequences that enable
// italic (cursive) text formatting in terminals that support ANSI styles.
//
// The formatting is reset after the string to ensure that subsequent
// terminal output is displayed using the default text style.
//
// Note: Not all terminal emulators support ANSI italic formatting.
// Unsupported terminals may ignore the escape sequence.
func Cursive(prompt string) string {
	// "\033[3m" enables italic text.
	// "\033[0m" resets all text formatting back to the terminal default.
	return "\033[3m" + prompt + "\033[0m"
}
