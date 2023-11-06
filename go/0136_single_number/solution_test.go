package single_number

import (
	"math"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{2, 2, 1}, want: 1},
		{input: []int{4, 1, 2, 1, 2}, want: 4},
		{input: []int{-8, 5, -3, 0, -8, 0, 5}, want: -3},
		{input: []int{math.MinInt, 1, math.MaxInt, 7, -2, 1, -2, math.MaxInt, 7}, want: math.MinInt},
		{input: []int{1}, want: 1},
	}

	for i, tc := range tests {
		got := singleNumber(tc.input)
		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
