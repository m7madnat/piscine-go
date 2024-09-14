package piscine

func FirstRune(s string) rune {
	res := []rune(s)
	return res[0]
}

// func FirstRune(s string) rune {
// 	res := []rune(s)
// 	return res[0]
// }

// func PointOne(n *int) {	*n = 1} 
// func UltimatePointOne(n ***int) {	***n = 1}
// func DivMod(a int, b int, div *int, mod *int) {	*div = a / b	*mod = a % b}
// func UltimateDivMod(a *int, b *int) {	y := *a	*a = *a / *b	*b = y % *b}
// func PrintStr(s string) {	for _, str := range s {		z01.PrintRune(str)	}}
// func StrLen(s string) int {	len := 0	for range s {		len++	}	return len}
// package piscine
// func FirstRune(s string) rune {
// 	res := []rune(s)
// 	return res[0]
// }
// package piscine
// func LastRune(s string) rune {
// 	res := []rune(s)
// 	return res[len(res)-1]
// }
// package piscine
// func NRune(s string, n int) rune {
// 	if n <= 0 || n > len(s) {
// 		return 0
// 	}
// 	res := []rune(s)
// 	return res[n-1]
// }
// package piscine
// func Compare(a, b string) int {
// 	if a < b {
// 		return -1
// 	} else if a > b {
// 		return 1
// 	} else {
// 		return 0
// 	}
// }
// package piscine
// func AlphaCount(s string) int {
// 	runes := []rune(s)
// 	count := 0
// 	for _, str := range runes {
// 		if str >= 'a' && str <= 'z' || str >= 'A' && str <= 'Z' {
// 			count++
// 		}
// 	}
// 	return count
// }
// package piscine
// func Index(s string, toFind string) int {
// 	length := 0
// 	sublength := 0
// 	for i := range s {
// 		length = i + 1
// 	}
// 	for i := range toFind {
// 		sublength = i + 1
// 	}
// 	t := 0
// 	for i := 0; i < length; i++ {
// 		j := 0
// 		t = i
// 		for j < sublength {
// 			if t < length && s[t] == toFind[j] {
// 				j++
// 				t++
// 			} else {
// 				break
// 			}
// 		}
// 		if j == sublength {
// 			return i
// 		}
// 	}
// 	return -1
// }
// package piscine
// func Concat(str1 string, str2 string) string {
// 	str3 := str1 + str2
// 	return str3
// }
// package piscine
// func IsUpper(str string) bool {
// 	for i := 0; i < len(str); i++ {
// 		if !(str[i] >= 'A' && str[i] <= 'Z') {
// 			return false
// 		}
// 	}
// 	return true
// }
// package piscine
// func IsLower(str string) bool {
// 	for i := 0; i < len(str); i++ {
// 		if !(str[i] >= 'a' && str[i] <= 'z') {
// 			return false
// 		}
// 	}
// 	return true
// }
// package piscine
// func IsAlpha(str string) bool {
// 	for i := 0; i < len(str); i++ {
// 		if !(str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z' || str[i] >= '0' && str[i] <= '9') {
// 			return false
// 		}
// 	}
// 	return true
// }
// package piscine
// func IsNumeric(str string) bool {
// 	for _, ch := range str {
// 		if ch < '0' || ch > '9' {
// 			return false
// 		}
// 	}
// 	return true
// }
// package piscine
// func checkprint(a rune) bool {
// 	if a >= 0 && a <= 31 {
// 		return true
// 	}
// 	return false
// }
// func IsPrintable(str string) bool {
// 	runes := []rune(str)
// 	for i := range runes {
// 		if checkprint(runes[i]) == true {
// 			return false
// 		}
// 	}
// 	return true
// }
// package piscine
// func ToUpper(s string) string {
// 	str := []rune(s)
// 	for i := 0; i < StrLen(s); i++ {
// 		if str[i] >= 97 && str[i] <= 122 {
// 			str[i] -= 32
// 		}
// 	}
// 	return string(str)
// }
// package piscine
// func ToLower(s string) string {
// 	str := make([]rune, len(s))
// 	for i, ch := range s {
// 		if ch >= 'A' && ch <= 'Z' {
// 			ch += 32
// 		}
// 		str[i] = ch
// 	}
// 	return string(str)
// }
// package main
// import (
// 	"os"
// 	"github.com/01-edu/z01"
// )
// func main() {
// 	arg := os.Args[0][2:]
// 	for _, r := range arg {
// 		z01.PrintRune(r)
// 	}
// 	z01.PrintRune('\n')
// }
// package main
// import (
// 	"os"
// 	"github.com/01-edu/z01"
// )
// func main() {
// 	arg := os.Args
// 	for i := range arg {
// 		if i != 0 {
// 			for j := range arg[i][:] {
// 				runes := []rune(arg[i])
// 				{
// 					z01.PrintRune(runes[j])
// 				}
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }
// package main
// import (
// 	"os"
// 	"github.com/01-edu/z01"
// )
// func main() {
// 	arg := os.Args
// 	for i := len(arg) - 1; i > 0; i-- {
// 		for _, c := range arg[i] {
// 			z01.PrintRune(c)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }