package input

import "fmt"

func ReadFloat(prompt string) float64 {
        fmt.Print(prompt)
        var userInput float64
        fmt.Scan(&userInput)
        return userInput
}
