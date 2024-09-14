package piscine

func AlphaCount(s string) int {
	runes := []rune(s)
	count := 0
	for _, str := range runes {
		if str >= 'a' && str <= 'z' || str >= 'A' && str <= 'Z' {
			count++
		}
	}
	return count
}
