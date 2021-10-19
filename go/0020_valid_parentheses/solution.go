package valid_parentheses

var parenthesesMapping = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	if s == "" {
		return true
	}

	stack := make([]byte, 0, len(s)) // Pseudo stack
	for _, r := range s {
		char := byte(r)
		if pairChar, found := parenthesesMapping[char]; found && len(stack) > 0 {
			lastChar := stack[len(stack)-1:][0] // Pop from stack
			stack = stack[:len(stack)-1]
			if lastChar != pairChar {
				return false
			}
			continue
		}
		stack = append(stack, char) // Push to stack
	}

	return len(stack) == 0
}
