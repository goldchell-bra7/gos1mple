package main

import (
	"fmt"
	"time"

	"github.com/goldchell-bra7/gos1mple/color"
	"github.com/goldchell-bra7/gos1mple/text"
	"github.com/goldchell-bra7/gos1mple/input"
)

func main() {
	fmt.Println(text.Header("Search files", "=", 40))

	question := input.ReadBool("Do you want to find a file?: ")

	if question {
		for i := 0; i <= 79; i++ {
			fmt.Println(color.Term("Searching files...", color.Green))
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println(color.Term("file found!", color.BrightYellow))
	} else {
		fmt.Println("Ok, next time")
	}
}
