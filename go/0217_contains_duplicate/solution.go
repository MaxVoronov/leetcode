package contains_duplicate

func containsDuplicate(nums []int) bool {
	storage := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, found := storage[num]; found {
			return true
		}

		storage[num] = struct{}{}
	}

	return false
}
