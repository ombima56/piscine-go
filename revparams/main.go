package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	r := len(args) - 1
	for r >= 0 {
		for _, q := range args[r] {
			z01.PrintRune(q)
		}
		r--
		z01.PrintRune('\n')
	}
}
