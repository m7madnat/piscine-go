package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := 1; i < len(args); i++ {
		for j := i; j > 0 && args[j-1] > args[j]; j-- {
			args[j], args[j-1] = args[j-1], args[j]
		}
	}

	for _, s := range args {
		for _, c := range s {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
