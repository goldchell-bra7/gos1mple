package main

import (
	"fmt"

	"github.com/goldchell-bra7/gos1mple/text"
)

func main() {
	fmt.Println(text.Bold("Hello"))
	fmt.Println(text.Cursive("Hello"))
	fmt.Println(text.Underline("Hello"))
	fmt.Println(text.Strike("Hello"))
}
