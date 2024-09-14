package piscine

func ToLower(s string) string {
	str := make([]rune, len(s))
	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			ch += 32
		}
		str[i] = ch
	}
	return string(str)
}
