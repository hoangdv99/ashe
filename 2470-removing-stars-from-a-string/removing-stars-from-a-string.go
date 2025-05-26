func removeStars(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c != '*' {
			stack = append(stack, c)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
