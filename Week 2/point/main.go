package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	Printout(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	Printout(points.y)
	z01.PrintRune('\n')
}

func Printout(num int) {
	var n []int
	var x int
	for num != 0 {
		x = num % 10
		n = append(n, x)
		num = num / 10
	}
	for i := len(n) - 1; i >= 0; i-- {
		z01.PrintRune(rune(n[i]) + 48)
	}
}
