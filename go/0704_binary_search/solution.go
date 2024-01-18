package binary_search

func Search(nums []int, target int) int {
	lowIdx, midIdx, highIdx := 0, 0, len(nums)-1

	for lowIdx <= highIdx {
		midIdx = lowIdx + (highIdx-lowIdx)/2
		if nums[midIdx] == target {
			return midIdx
		}

		if nums[midIdx] < target {
			lowIdx = midIdx + 1
		} else {
			highIdx = midIdx - 1
		}
	}

	return -1
}
