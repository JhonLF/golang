package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	fmt.Println(arguments)
	r := []rune(arguments)
	for _, v := range r[2:] {
		z01.PrintRune(v)
	}

	z01.PrintRune('\n')
}
