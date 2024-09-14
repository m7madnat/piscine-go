package piscine

func IsAlpha(str string) bool {
	for i := 0; i < len(str); i++ {
		if !(str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z' || str[i] >= '0' && str[i] <= '9') {
			return false
		}
	}
	return true
}
