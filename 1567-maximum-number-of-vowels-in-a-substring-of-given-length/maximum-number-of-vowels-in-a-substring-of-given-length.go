func maxVowels(s string, k int) int {
	var windowCount int
	for i := range k {
		if isVowel(s[i]) {
			windowCount++
		}
	}

    max := windowCount

	for i := 1; i <= len(s)-k; i++ {
		if !isVowel(s[i-1]) && isVowel(s[i+k-1]) {
			windowCount++
		} else if isVowel(s[i-1]) && !isVowel(s[i+k-1]) {
			windowCount--
		}
		if windowCount > max {
			max = windowCount
		}
	}

	return max
}

func isVowel(b byte) bool {
	if b == 'a' || b == 'u' || b == 'e' || b == 'i' || b == 'o' {
		return true
	}
	return false
}