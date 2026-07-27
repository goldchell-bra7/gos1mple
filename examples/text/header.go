package main

import (
	"fmt"

	"github.com/goldchell-bra7/gos1mple/input"
	"github.com/goldchell-bra7/gos1mple/text"
)

func main() {
	header := text.Header("App for you", "=", 20)
	fmt.Println(header)

	question := input.ReadBool("Do you love me?: ")

	if question {
		fmt.Println("I love you too!")
	} else {
		fmt.Println(" :( ")
	}
}
