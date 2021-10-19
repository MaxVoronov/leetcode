package move_zeroes

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		}
	}

	if idx > 0 {
		for ; idx < len(nums); idx++ {
			nums[idx] = 0
		}
	}
}
