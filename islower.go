package piscine

func IsLower(str string) bool {
	for i := 0; i < len(str); i++ {
		if !(str[i] >= 'a' && str[i] <= 'z') {
			return false
		}
	}
	return true
}
