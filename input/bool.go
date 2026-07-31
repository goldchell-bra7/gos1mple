package input

import "fmt"

// ReadBool displays a prompt, reads a boolean value from standard input,
// and returns the entered result.
//
// This helper function provides a reusable way to read boolean input
// (typically "true" or "false") from the console, keeping input logic
// consistent across the project.
func ReadBool(prompt string) bool {
	// Display the prompt without appending a newline.
	fmt.Print(prompt)

	// Store the boolean value entered by the user.
	var userInput bool

	// Read a single boolean value from standard input.
	fmt.Scan(&userInput)

	// Return the parsed boolean value.
	return userInput
}
