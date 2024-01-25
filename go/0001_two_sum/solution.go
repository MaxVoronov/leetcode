package two_sum

func twoSum(nums []int, target int) []int {
	usedNums := make(map[int]int, len(nums))
	for idx, num := range nums {
		if diffIdx, found := usedNums[target-num]; found {
			return []int{diffIdx, idx}
		}

		usedNums[num] = idx
	}

	return []int{}
}
