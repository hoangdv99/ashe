func isPalindrome(x int) bool {
	stringValue := strconv.Itoa(x)
	reversedString := reverseString((stringValue))
	return stringValue == reversedString
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}