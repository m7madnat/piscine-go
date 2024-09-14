package piscine

func ToUpper(s string) string {
	str := []rune(s)

	for i := 0; i < StrLen(s); i++ {
		if str[i] >= 97 && str[i] <= 122 {
			str[i] -= 32
		}
	}

	return string(str)
}

// func ToUpper(s string) string {
// 	result := ""
// 	for _, char := range s {
// 		if char >= 'a' && char <= 'z' {
// 			result += string(char - ('a' - 'A'))
// 		} else {
// 			result += string(char)
// 		}
// 	}
// 	return result
// }
