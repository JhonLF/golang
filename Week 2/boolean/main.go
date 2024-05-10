package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return true
	} else {
		return false
	}
}

//func lenArg(x []string) int {
//	count := 0
//	for i := 0; i < len(x); i++ {
//		for _, v := range x[i] {
//			if v != ' ' {
//				count += 1
//			}
//		}
//	}
//	return count
//}

func main() {
	if isEven(len(os.Args)) == true {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
