func strStr(haystack string, needle string) int {
    if haystack == needle {
        return 0
    }
	for i := range haystack {
		if i+len(needle) <= len(haystack) && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
