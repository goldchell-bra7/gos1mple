package main

import (
    "fmt"

    "github.com/goldchell-bra7/gos1mple/input"
)

func main() {
    question := input.ReadBool("Do you like Go?: ")

    if question {
        fmt.Println("Great!")
    } else {
        fmt.Println("Maybe someday")
    }
}
