
func isPalindrome(s string) bool {
	formattedString := strings.ToLower(removeNonAlphanumeric(s))
	result := true
	for i := 0; i < len(formattedString)/2; i++ {
		if formattedString[i] != formattedString[len(formattedString)-1-i] {
			result = false
			break
		}
	}
	return result
}

func removeNonAlphanumeric(s string) string {
	var builder strings.Builder

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}