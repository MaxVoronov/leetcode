package palindrome_number

import "math"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	numLen := math.Floor(math.Log10(float64(x))) + 1
	halfLen := math.Floor(numLen / 2)
	for i := 0.; i < halfLen; i++ {
		if x/int(math.Pow(10, i))%10 != x/int(math.Pow(10, numLen-i-1))%10 {
			return false
		}
	}

	return true
}
