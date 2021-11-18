package sqrtx

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		midSq := mid * mid
		if midSq == x {
			return mid
		}

		if midSq < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
