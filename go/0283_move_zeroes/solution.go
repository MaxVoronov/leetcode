package move_zeroes

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	var wPos int
	for rPos := 0; rPos < len(nums); rPos++ {
		if nums[rPos] != 0 {
			nums[wPos], nums[rPos] = nums[rPos], nums[wPos]
			wPos++
		}
	}
}
