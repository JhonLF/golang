package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 1 {
		fmt.Println("File name missing")
	} else if len(arg) == 1 {
		f, err := os.ReadFile("quest8.txt")
		if err != nil {
			fmt.Println("error read file")
		} else {
			fmt.Print(string(f))
		}
	} else if len(arg) >= 2 {
		fmt.Println("Too many arguments")
	}
}
