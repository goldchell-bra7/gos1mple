package input

import "fmt"

// ReadStringln displays a prompt followed by a newline, reads a string
// from standard input, and returns the entered value.
//
// Unlike ReadString, this function prints the prompt on a separate line,
// making console interactions more readable. Like fmt.Scan, it reads
// input only until the first whitespace character. To read an entire
// line containing spaces, use a buffered reader instead.
func ReadStringln(prompt string) string {
	// Print the prompt and move the cursor to the next line.
	fmt.Printf("%s\n", prompt)

	// Store the string entered by the user.
	var userInput string

	// Read a single whitespace-delimited string from standard input.
	fmt.Scan(&userInput)

	// Return the entered string.
	return userInput
}
