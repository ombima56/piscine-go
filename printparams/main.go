package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]
	for _, args := range name {
		for _, ch := range args {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
