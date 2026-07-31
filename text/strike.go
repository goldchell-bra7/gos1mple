package text

// Strike wraps the provided string with ANSI escape sequences that apply
// a strikethrough effect in terminals that support ANSI text formatting.
//
// The formatting is reset after the string to ensure that subsequent
// terminal output is displayed normally.
//
// Note: Strikethrough is not supported by all terminal emulators.
// Unsupported terminals may ignore the escape sequence.
func Strike(prompt string) string {
	// "\033[9m" enables strikethrough text.
	// "\033[0m" resets all text formatting back to the terminal default.
	return "\033[9m" + prompt + "\033[0m"
}
