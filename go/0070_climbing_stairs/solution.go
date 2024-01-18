package climbing_stairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prevSteps, currSteps := 1, 2 // 1 stair ways, 2 stairs ways
	for i := 3; i <= n; i++ {
		currSteps, prevSteps = prevSteps+currSteps, currSteps
	}

	return currSteps
}
