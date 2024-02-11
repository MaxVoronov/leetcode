package maximum_subarray

func maxSubArray(nums []int) int {
	maxSum, currSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(currSum+nums[i], nums[i])
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}
