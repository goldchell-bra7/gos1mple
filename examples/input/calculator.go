package main

import (
    "fmt"

    "github.com/goldchell-bra7/gos1mple/input"
)

func main() {
    first := input.ReadFloat("Enter the first number: ")
    second := input.ReadFloat("Enter the second number: ")
    operator := input.ReadString("Select an operation (+,-): ")

    var result float64

    switch operator {
    case "+":
        result = first + second
        fmt.Println("Result:", result)

    case "-":
        result = first - second
        fmt.Println("Result:", result)

    default:
        fmt.Println("Wrong operator")
    }
}
