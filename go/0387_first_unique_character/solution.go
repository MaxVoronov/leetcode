package first_unique_character

// Use simple array of structs
func firstUniqChar(s string) int {
	const alphabetLen = 26
	type freqData struct {
		amount, firstIdx int
	}

	var charsFreq [alphabetLen]freqData

	for idx := 0; idx < len(s); idx++ {
		data := &charsFreq[s[idx]-'a']
		if data.firstIdx == 0 {
			data.firstIdx = idx
		}
		data.amount++
	}

	for idx := 0; idx < len(s); idx++ {
		data := &charsFreq[s[idx]-'a']
		if data.amount == 1 {
			return data.firstIdx
		}
	}

	return -1
}

// Use maps
func firstUniqCharMaps(s string) int {
	const alphabetLen = 26
	charsFreq := make(map[byte]uint32, alphabetLen)

	for i := 0; i < len(s); i++ {
		charsFreq[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if charsFreq[s[i]] == 1 {
			return i
		}
	}

	return -1
}
