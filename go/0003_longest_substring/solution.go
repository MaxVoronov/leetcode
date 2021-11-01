package longest_substring

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var maxLen int
	r := []rune(s)
	usedChars := make(map[rune]struct{})

	for i := 0; i < len(r); i++ {
		if _, found := usedChars[r[i]]; found {
			i -= len(usedChars)
			usedChars = make(map[rune]struct{})
			continue
		}

		usedChars[r[i]] = struct{}{}
		if maxLen < len(usedChars) {
			maxLen = len(usedChars)
		}
	}

	return maxLen
}
