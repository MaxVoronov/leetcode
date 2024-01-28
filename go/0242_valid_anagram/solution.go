package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]int16, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}

	for _, val := range m {
		if val != 0 {
			return false
		}
	}

	return true
}
