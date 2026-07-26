package input

import "fmt"

func ReadBoolln(prompt string) bool {
        fmt.Printf("%s\n", prompt)
        var userInput bool
        fmt.Scan(&userInput)
        return userInput
}
