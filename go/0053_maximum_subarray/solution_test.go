package maximum_subarray

import "testing"

func Test_maxSubArray(t *testing.T) {
	var testCases = []struct {
		input []int
		want  int
	}{
		{input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6}, // [4, -1, 2, 1]
		{input: []int{5, 4, -1, 7, 8}, want: 23},               // [5, 4, -1, 7, 8]
		{input: []int{1}, want: 1},                             // [1]
	}

	for i, tc := range testCases {
		got := maxSubArray(tc.input)
		if got != tc.want {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
