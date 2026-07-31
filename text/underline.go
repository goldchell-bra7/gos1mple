package text

// Underline wraps the provided string with ANSI escape sequences that
// apply underline formatting in terminals that support ANSI text styles.
//
// The formatting is reset after the string to ensure that subsequent
// terminal output is displayed using the default style.
//
// Note: Underline support depends on the terminal emulator. Unsupported
// terminals may ignore the escape sequence.
func Underline(prompt string) string {
	// "\033[4m" enables underlined text.
	// "\033[0m" resets all text formatting back to the terminal default.
	return "\033[4m" + prompt + "\033[0m"
}
