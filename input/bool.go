package input

import "fmt"

func ReadBool(prompt string) bool {
	fmt.Print(prompt)
	var userInput bool
	fmt.Scan(&userInput)
	return userInput
}
