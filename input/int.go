package input

import "fmt"

// ReadInt displays a prompt, reads an integer from standard input,
// and returns the entered value.
//
// This helper function centralizes integer input handling, allowing
// the same logic to be reused throughout the project instead of
// duplicating input code in multiple places.
func ReadInt(prompt string) int {
	fmt.Print(prompt)

	// Store the integer value entered by the user.
	var userInput int

	// Read a single integer from standard input.
	// The entered value is written directly into userInput.
	fmt.Scan(&userInput)

	// Return the parsed integer to the caller.
	return userInput
}
