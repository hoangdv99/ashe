func longestCommonPrefix(strs []string) string {
	var result string
	for i, v := range strs[0] {
		duplicated := false
		for _, str := range strs {
			duplicated = i <= len(str)-1 && string(str[i]) == string(v)
			if !duplicated {
				break
			}
		}
		if duplicated {
			result += string(v)
		} else {
            break
        }
	}
	return result
}