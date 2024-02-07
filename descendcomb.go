package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for l := '9'; l >= '0'; l-- {
		for n := '9'; n >= '0'; n-- {
			for r := '9'; r >= '0'; r-- {
				for s := '9'; s >= '0'; s-- {
					if l > r || (l == r && n > s) {
						z01.PrintRune(l)
						z01.PrintRune(n)
						z01.PrintRune(' ')
						z01.PrintRune(r)
						z01.PrintRune(s)
						if !(l == '0' && n == '1' && r == '0' && s == '0') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
