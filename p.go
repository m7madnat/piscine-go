package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	// PrintComb()
	// fmt.Println(strrev("alooo"))
	// fmt.Println(BasicAtoi("001234"))
	// fmt.Println(Fibonacci(-5))
	// fmt.Println(FirstRune("abc"))
}

func StrRev(s string) string {
	lenStr := len(s) - 1
	var indexInAnewArr int
	newStr := make([]rune, len(s))
	for i := lenStr; i >= 0; i-- {
		newStr[indexInAnewArr] = rune(s[i])
		indexInAnewArr++
	}
	return string(newStr)
}

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					if (a*10 + b) < (c*10 + d) {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(' ')
						z01.PrintRune(c)
						z01.PrintRune(d)
						if !(a == '9' && b == '8') {
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

func strrev(s string) string {
	var res string
	for _, i := range s {
		res = string(i) + res
	}
	return res
}

func UltimateDivMod(a *int, b *int) {
	mod := *a % *b // mod % div /
	div := *a / *b
	*a, *b = div, mod
}

func BasicAtoi(s string) int {
	number := 0
	c := 0
	for _, word := range s {
		for i := '0'; i < word; i++ {
			c++
		}
		number = number*10 + c
		c = 0
	}
	return number
}

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return RecursiveFactorial(nb-1) * nb
}

func IterativePower(nb int, power int) int {
	if power <= 0 || power > 20 {
		return 0
	}
	res := 1
	for i := 1; i <= power; i++ {
		res *= nb
	}
	return res
}

func RecursivePower(nb int, power int) int {
	if power < 0 || power > 20 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return nb * RecursivePower(nb, power-1)
}

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index < 2 {
		return index
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}