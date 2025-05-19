func isSubsequence(s string, t string) bool {
    if s == "" {
        return true
    }
	runesT := []rune(s)
	var j int
	for _, c := range t {
		if c == runesT[j] {
			j++
		}
        if j == len(s) {
			return true
		}
	}
	return false
}