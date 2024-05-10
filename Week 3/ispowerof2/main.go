package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) == 1 {
		arr := os.Args[1]
		num, _ := strconv.Atoi(arr)
		if Ipower(num) == true {
			z01.PrintRune('t')
			z01.PrintRune('r')
			z01.PrintRune('u')
			z01.PrintRune('e')
			z01.PrintRune('\n')
		} else {
			z01.PrintRune('f')
			z01.PrintRune('a')
			z01.PrintRune('l')
			z01.PrintRune('s')
			z01.PrintRune('e')
			z01.PrintRune('\n')
		}
	}
}

func Ipower(n int) bool {
	x := 2
	if n == 2 || n == 1 {
		return true
	}
	for i := 2; i < n; i++ {
		for j := 0; j < i; j++ {
			x *= 2
			if x == n {
				return true
			}
		}
	}
	return false
}
