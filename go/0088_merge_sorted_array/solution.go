package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1], nums1[m-1] = nums1[m-1], 0
			m--
			continue
		}

		nums1[m+n-1] = nums2[n-1]
		n--
	}

	if n > 0 {
		for n > 0 {
			nums1[n-1] = nums2[n-1]
			n--
		}
	}
}
