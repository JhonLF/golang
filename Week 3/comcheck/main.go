package main

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]

	for _, v := range arr {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
