package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	valor_string := fmt.Sprint(n)
	for _, x := range valor_string {
		z01.PrintRune(x)
	}
}
