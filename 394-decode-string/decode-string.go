func decodeString(s string) string {
	var countStack []int
	var stringStack []string
	currStr := ""
	num := 0

	for _, c := range s {
		switch {
		case unicode.IsDigit(c):
			num = num*10 + int(c-'0')

		case c == '[':
			countStack = append(countStack, num)
			stringStack = append(stringStack, currStr)
			num = 0
			currStr = ""

		case c == ']':
			repeat := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			prevStr := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]

			currStr = prevStr + strings.Repeat(currStr, repeat)

		default:
			currStr += string(c)
		}
	}

	return currStr
}
