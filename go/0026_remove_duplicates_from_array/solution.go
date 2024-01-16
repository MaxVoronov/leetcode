package remove_duplicates_from_array

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	lastNum := nums[0] // Last number for comparing
	rCur, wCur := 1, 1 // Read and write cursors

	for rCur < len(nums) {
		if nums[rCur] > lastNum {
			lastNum = nums[rCur]
			nums[wCur] = lastNum
			wCur++
		}

		rCur++
	}

	return wCur
}
