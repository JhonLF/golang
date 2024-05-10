package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	sort(arguments)
}

func sort(s []string) {
	for i := 0; i < len(s)-1; i++ {
		for x := 0; x < len(s)-i-1; x++ {
			if s[x] > s[x+1] {
				s[x], s[x+1] = s[x+1], s[x]
			}
		}
	}
}

func printout(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
