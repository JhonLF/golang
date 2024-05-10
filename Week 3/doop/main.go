package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1:]
	if len(arr) != 3 {
		return
	}

	n1, err1 := strconv.Atoi(arr[0])
	n2, err2 := strconv.Atoi(arr[2])
	if err1 != nil || err2 != nil || overflowCheck(n1) || overflowCheck(n2) {
		return
	}

	switch arr[1] {
	case "+":
		r := n1 + n2
		printRune(r)
		z01.PrintRune('\n')
	case "*":
		r := n1 * n2
		printRune(r)
		z01.PrintRune('\n')
	case "-":
		r := n1 - n2
		printRune(r)
		z01.PrintRune('\n')
	case "/":
		if n2 == 0 {
			printStr("No division by 0")
			z01.PrintRune('\n')
		} else {
			r := n1 / n2
			printRune(r)
			z01.PrintRune('\n')
		}
	case "%":
		if n2 == 0 {
			printStr("No modulo by 0")
			z01.PrintRune('\n')
		} else {
			r := n1 % n2
			printRune(r)
			z01.PrintRune('\n')
		}
	}
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func printRune(num int) {
	if num == 0 {
		z01.PrintRune('0')
		return
	}

	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}

	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune('0' + digits[i]))
	}
}

func overflowCheck(num int) bool {
	return num > 1000000000 || num < -1000000000
}
