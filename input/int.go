package input

import "fmt"

func ReadInt(prompt string) int {
	fmt.Print(prompt)
	var userInput int
	fmt.Scan(&userInput)
	return userInput
}
