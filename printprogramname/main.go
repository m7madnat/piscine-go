package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0][2:]

	for _, r := range arg {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
