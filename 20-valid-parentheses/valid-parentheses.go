func isValid(s string) bool {
	// Use a map with rune keys and values for better performance
	symbolMap := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	stack := []rune{} // Use a rune slice for the stack

	for _, ch := range s {
		if closing, found := symbolMap[ch]; found { // Check if it's an opening bracket
			stack = append(stack, closing) // Append the corresponding closing bracket to the stack
		} else { // It's a closing bracket
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1] // Pop from the stack
		}
	}

	return len(stack) == 0
}