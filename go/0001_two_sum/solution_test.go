package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   []int
	}{
		{input: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{input: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{input: []int{0, 4, 3, 0}, target: 0, want: []int{0, 3}},
		{input: []int{-1, -2, -3, -4, -5}, target: -8, want: []int{2, 4}},
	}

	for i, tc := range tests {
		got := twoSum(tc.input, tc.target)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
