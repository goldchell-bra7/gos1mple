package input

import "fmt"

// ReadFloat displays a prompt, reads a floating-point number from
// standard input, and returns the entered value.
//
// The function always returns a float64, making it suitable for
// calculations that require decimal precision. Centralizing this
// logic also helps keep user input handling consistent across the project.
func ReadFloat(prompt string) float64 {
	// Display the prompt without appending a newline.
	fmt.Print(prompt)

	// Store the floating-point value entered by the user.
	var userInput float64

	// Read a single floating-point value from standard input.
	fmt.Scan(&userInput)

	// Return the parsed floating-point value.
	return userInput
}
