package input

import "fmt"

func ReadString(prompt string) string {
	fmt.Print(prompt)
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}
