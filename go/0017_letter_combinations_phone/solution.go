package letter_combinations_phone

var digitsMapping = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	return combiner(digits, []string{})
}

func combiner(digits string, combs []string) []string {
	if digits == "" {
		return combs
	}

	letters := digitsMapping[string(digits[0])]
	newCombs := make([]string, 0, len(letters))

	if len(combs) == 0 {
		newCombs = letters
	} else {
		for _, c := range combs {
			for _, l := range letters {
				newCombs = append(newCombs, c+l)
			}
		}
	}

	return combiner(digits[1:], newCombs)
}
