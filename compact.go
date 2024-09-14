package piscine

func Compact(ptr *[]string) int {
	src := *ptr
	dst := make([]string, 0, len(src))
	for _, s := range src {
		if s != "" {
			dst = append(dst, s)
		}
	}
	*ptr = dst
	return len(dst)
}
