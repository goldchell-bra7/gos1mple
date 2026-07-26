package input

import "fmt"

func ReadStringln(prompt string) string {
        fmt.Printf("%s\n", prompt)
        var userInput string
        fmt.Scan(&userInput)
        return userInput
}
