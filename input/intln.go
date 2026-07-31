package input

import "fmt"

// ReadIntln displays a prompt followed by a newline, reads an integer
// from standard input, and returns the entered value.
//
// Unlike ReadInt, this function prints the prompt on its own line,
// which can improve readability for multi-line console interactions.
func ReadIntln(prompt string) int {
	// Print the prompt and move the cursor to the next line.
	fmt.Printf("%s\n", prompt)

	// Store the integer value entered by the user.
	var userInput int

	// Read a single integer from standard input.
	fmt.Scan(&userInput)

	// Return the parsed integer to the caller.
	return userInput
}

