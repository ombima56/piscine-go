package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printNum(num int) {
	z := '0'
	for i := 0; i < num; i++ {
		z++
	}
	z01.PrintRune(z)
}

func printStr(s string) {
	for _, arg := range s {
		z01.PrintRune(arg)
	}
}

func main() {
	part1 := "x = "
	part2 := ", y = "
	points := &point{}
	setPoint(points)

	printStr(part1)
	printNum(points.x / 10)
	printNum(points.x % 10)
	printStr(part2)
	printNum(points.y / 10)
	printNum(points.y % 10)
	z01.PrintRune('\n')
}
