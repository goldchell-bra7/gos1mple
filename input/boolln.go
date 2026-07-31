package input

import "fmt"

// ReadBoolln displays a prompt followed by a newline, reads a boolean
// value from standard input, and returns the entered result.
//
// Unlike ReadBool, this function prints the prompt on a separate line,
// which can make console output easier to read in interactive applications.
func ReadBoolln(prompt string) bool {
	// Print the prompt and move the cursor to the next line.
	fmt.Printf("%s\n", prompt)

	// Store the boolean value entered by the user.
	var userInput bool

	// Read a single boolean value from standard input.
	fmt.Scan(&userInput)

	// Return the parsed boolean value.
	return userInput
}
