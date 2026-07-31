package input

import "fmt"

// ReadFloatln displays a prompt followed by a newline, reads a
// floating-point number from standard input, and returns the entered value.
//
// Unlike ReadFloat, this function places the prompt on its own line,
// making console interactions cleaner when longer or multi-line prompts
// are used.
func ReadFloatln(prompt string) float64 {
	// Print the prompt and move the cursor to the next line.
	fmt.Printf("%s\n", prompt)

	// Store the floating-point value entered by the user.
	var userInput float64

	// Read a single floating-point value from standard input.
	fmt.Scan(&userInput)

	// Return the parsed floating-point value.
	return userInput
}
