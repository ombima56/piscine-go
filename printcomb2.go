package main

import "github.com/01-edu/z01"

func PrintComb2() {
	for l := '0'; l <= '9'; l++ {
		for n := '0'; n <= '9'; n++ {
			for r := '0'; r <= '9'; r++ {
				for s := '0'; s <= '9'; s++ {
					if l < r || (l == r && n < s) {
						z01.PrintRune(l)
						z01.PrintRune(n)
						z01.PrintRune(' ')
						z01.PrintRune(r)
						z01.PrintRune(s)
						if l != '9' || n != '9' || r != '8' || s != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintComb2()
}
