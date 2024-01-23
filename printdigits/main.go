package main

import "github.com/01-edu/z01"

func main() {
	for j := '0'; j <= '9'; j++ {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}
