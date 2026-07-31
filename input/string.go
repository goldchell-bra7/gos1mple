package input

import "fmt"

// ReadString displays a prompt, reads a string from standard input,
// and returns the entered value.
//
// This function uses fmt.Scan, which reads input until the first
// whitespace character. It is intended for single-word input.
// For reading an entire line (including spaces), a buffered reader
// should be used instead.
func ReadString(prompt string) string {
	// Display the prompt without appending a newline.
	fmt.Print(prompt)

	// Store the string entered by the user.
	var userInput string

	// Read a single whitespace-delimited string from standard input.
	fmt.Scan(&userInput)

	// Return the entered string.
	return userInput
}
