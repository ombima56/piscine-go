package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	for _, char := range name[2:] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
