package input

import "fmt"

func ReadFloatln(prompt string) float64 {
        fmt.Printf("%s\n", prompt)
        var userInput float64
        fmt.Scan(&userInput)
        return userInput
}
