package input

import "fmt"

func ReadIntln(prompt string) int {
        fmt.Printf("%s\n", prompt)
        var userInput int
        fmt.Scan(&userInput)
        return userInput
}

