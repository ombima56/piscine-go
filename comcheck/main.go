package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for _, ch := range arg {
		if ch == "01" || ch == "galaxy" || ch == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
