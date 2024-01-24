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
		{input: []int{2, 0, 8, 5, 0, 0}, want: []int{2, 8, 5, 0, 0, 0}},
		{
			input: []int{73348, 66106, -85359, 53996, 18849, -6590, -53381, -86613, -41065, 83457, 0},
			want:  []int{73348, 66106, -85359, 53996, 18849, -6590, -53381, -86613, -41065, 83457, 0},
		},
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
