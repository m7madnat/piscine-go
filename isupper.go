package piscine

func IsUpper(str string) bool {
	for i := 0; i < len(str); i++ {
		if !(str[i] >= 'A' && str[i] <= 'Z') {
			return false
		}
	}
	return true
}
