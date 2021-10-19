package move_zeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{7, 5, 3, 1}, want: []int{7, 5, 3, 1}},
		{input: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		{input: []int{0, 0, 1}, want: []int{1, 0, 0}},
		{input: []int{0}, want: []int{0}},
		{input: []int{}, want: []int{}},
	}

	for i, tc := range tests {
		moveZeroes(tc.input)
		if !reflect.DeepEqual(tc.want, tc.input) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, tc.input)
		}
	}
}
