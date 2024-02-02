package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args
	count := 0

	for r := range s {
		count = r + 1
	}
	for i := 1; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	for j := 1; j <= count-1; j++ {
		for _, el := range s[j] {
			z01.PrintRune(el)
		}
		z01.PrintRune('\n')
	}
}
