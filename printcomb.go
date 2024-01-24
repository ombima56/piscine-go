package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for q := '0'; q <= '9'; q++ {
		for r := q + 1; r <= '9'; r++ {
			for s := r + 1; s <= '9'; s++ {
				z01.PrintRune(q)
				z01.PrintRune(r)
				z01.PrintRune(s)
				if q != '7' || r != '8' || s != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
