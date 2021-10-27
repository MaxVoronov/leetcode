package sort_colors

func sortColors(nums []int) {
	numLen := len(nums)
	if numLen <= 1 {
		return
	}

	// Swap elements in-place loop
	lShift, rShift := 0, 0
	for i := 0; i < numLen-rShift; i++ {
		if nums[i] == 0 {
			nums[i], nums[lShift] = nums[lShift], nums[i]
			lShift++
		} else if nums[i] == 2 {
			nums[i], nums[numLen-rShift-1] = nums[numLen-rShift-1], nums[i]
			rShift++
			i--
		}
	}
}
