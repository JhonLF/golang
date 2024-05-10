package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var revert []int
	if n == 0 {
		revert = append(revert, 0)
	}
	for n != 0 {
		revert = append(revert, n%10)
		n /= 10
	}
	for i := 0; i < len(revert)-1; i++ {
		for x := 0; x < len(revert)-i-1; x++ {
			if revert[x] > revert[x+1] {
				revert[x], revert[x+1] = revert[x+1], revert[x]
			}
		}
	}
	for _, v := range revert {
		z01.PrintRune(rune(v + 48))
	}
}
