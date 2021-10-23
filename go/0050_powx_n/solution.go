package powx_n

func myPow(x float64, n int) float64 {
	isEven := n%2 == 0

	// Optimization for special cases
	if n == 0 {
		return 1
	} else if n == 1 || x == 1 || (x == -1 && !isEven) {
		return x
	} else if x == -1 && isEven {
		return -x
	}

	isNegative := n < 0
	if isNegative {
		n *= -1
	}

	n /= 2
	result := 1.
	for ; n > 0; n-- {
		result *= x
	}

	result *= result
	if !isEven {
		result *= x
	}

	if isNegative {
		return 1 / result
	}

	return result
}
