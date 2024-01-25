package plus_one

func plusOne(digits []int) []int {
	result := make([]int, len(digits))
	extra := 1 // plus one or remains

	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i] + extra
		if num > 9 {
			result[i] = 0
			extra = 1
		} else {
			result[i] = num
			extra = 0
		}
	}

	if extra > 0 {
		result = append([]int{1}, result...)
	}

	return result
}
