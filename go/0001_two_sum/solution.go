package two_sum

func TwoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)

	for i, iVal := range nums {
		for j, jVal := range nums {
			if i != j && iVal+jVal == target {
				result = []int{i, j}
				return result
			}
		}
	}

	return result
}
