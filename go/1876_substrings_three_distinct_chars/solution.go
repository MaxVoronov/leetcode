package substrings_three_distinct_chars

func countGoodSubstrings(s string) int {
	const k = 3 // sliding window size

	runes := []rune(s)
	charsMemo := make(map[rune]int, k)
	result, leftIdx := 0, 0

	for rightIdx := 0; rightIdx < len(runes); rightIdx++ {
		charsMemo[runes[rightIdx]]++

		// Moving left side of sliding window and remove old character
		if rightIdx-leftIdx == k {
			oldChar := runes[leftIdx]

			charsMemo[oldChar]--
			if charsMemo[oldChar] == 0 {
				delete(charsMemo, oldChar) // clean chars with zero counters
			}
			leftIdx++
		}

		// If we contain k chars in memo, then all these chars are unique
		if len(charsMemo) == k {
			result++
		}
	}

	return result
}
