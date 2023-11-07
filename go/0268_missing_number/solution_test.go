package missing_number

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{3, 0, 1}, want: 2},
		{input: []int{0, 1}, want: 2},
		{input: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, want: 8},
		{input: []int{1}, want: 0},
	}

	for i, tc := range tests {
		got := missingNumber(tc.input)
		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
