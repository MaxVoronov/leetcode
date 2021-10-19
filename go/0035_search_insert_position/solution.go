package search_insert_position

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 || nums[0] > target {
		return 0
	}

	if nums[l-1] < target {
		return l
	}

	// Based on Binary search
	lowIdx, midIdx, highIdx := 0, 0, l-1
	for lowIdx <= highIdx {
		midIdx = (highIdx + lowIdx) / 2
		if nums[midIdx] == target {
			return midIdx
		}

		if nums[midIdx] < target {
			lowIdx = midIdx + 1
		} else {
			highIdx = midIdx - 1
		}
	}

	if nums[midIdx] <= target {
		midIdx++
	}

	return midIdx
}
