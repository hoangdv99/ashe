func strStr(haystack string, needle string) int {
	for i := range haystack {
		if i+len(needle) <= len(haystack) && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
