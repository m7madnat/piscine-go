package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func setPoint(ptr *point, x, y int) {
	ptr.x = x
	ptr.y = y
}

func main() {
	points := point{}
	setPoint(&points, 42, 21)

	printString("x = ")
	a(52)
	a(50)
	printString(", y = ")
	a(50)
	a(49)
	z01.PrintRune('\n')
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func a(r rune) {
	z01.PrintRune(r)
}
