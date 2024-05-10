package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	for _, v := range a {
		printout(v)
		z01.PrintRune('\n')
	}
}

func printout(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
