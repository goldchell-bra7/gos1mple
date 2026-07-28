package main

import (
	"fmt"

	"github.com/goldchell-bra7/gos1mple/color"
)

func main() {
	fmt.Println(color.Term("Black", color.Black))
	fmt.Println(color.Term("Red", color.Red))
	fmt.Println(color.Term("Green", color.Green))
	fmt.Println(color.Term("Yellow", color.Yellow))
	fmt.Println(color.Term("Blue", color.Blue))
	fmt.Println(color.Term("Magenta", color.Magenta))
	fmt.Println(color.Term("Cyan", color.Cyan))
	fmt.Println(color.Term("White", color.White))

}
