package missing_number

func missingNumber(nums []int) (result int) {
	for i := 0; i <= len(nums); i++ {
		result ^= i
	}

	for _, num := range nums {
		result ^= num
	}

	return
}
