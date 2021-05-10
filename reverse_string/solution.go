package reverse_string

func reverseString(s []byte) {
	l := len(s)
	if l < 2 {
		return
	}

	idx := 0
	for {
		pairIdx := l - idx - 1
		if pairIdx <= idx {
			break
		}

		s[idx], s[pairIdx] = s[pairIdx], s[idx]
		idx++
	}
}
