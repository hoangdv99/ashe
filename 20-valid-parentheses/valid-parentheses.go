func isValid(s string) bool {
	symbolMap := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}
	stack := []string{}

	for _, ch := range s {
		if string(ch) == "{" || string(ch) == "[" || string(ch) == "(" {
			stack = append([]string{string(ch)}, stack...)
		} else {
			if len(stack) > 0 && symbolMap[stack[0]] == string(ch) {
				stack = stack[1:]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}