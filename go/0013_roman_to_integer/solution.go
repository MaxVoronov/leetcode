package roman_to_integer

var numsMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	result, sLen := 0, len(s)
	for i := 0; i < sLen; i++ {
		num := numsMap[s[i]]
		if i > 0 {
			prevNum := numsMap[s[i-1]]
			if prevNum < num && (num/prevNum == 5 || num/prevNum == 10) {
				result += num - prevNum - prevNum
				continue
			}
		}
		result += num
	}

	return result
}
