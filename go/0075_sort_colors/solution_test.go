package sort_colors

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{2, 0, 2, 1, 1, 0}, want: []int{0, 0, 1, 1, 2, 2}},
		{input: []int{2, 0, 2, 1, 1, 1}, want: []int{0, 1, 1, 1, 2, 2}},
		{input: []int{2, 2, 2, 1, 1, 1}, want: []int{1, 1, 1, 2, 2, 2}},
		{input: []int{1, 2, 2, 0, 1}, want: []int{0, 1, 1, 2, 2}},
		{input: []int{0, 1, 2, 2}, want: []int{0, 1, 2, 2}},
		{input: []int{1, 2, 0}, want: []int{0, 1, 2}},
		{input: []int{2, 0, 1}, want: []int{0, 1, 2}},
		{input: []int{2, 2, 2}, want: []int{2, 2, 2}},
		{input: []int{1}, want: []int{1}},
		{input: []int{0}, want: []int{0}},
		{input: []int{}, want: []int{}},
	}

	for i, tc := range tests {
		sortColors(tc.input)
		if !reflect.DeepEqual(tc.want, tc.input) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, tc.input)
		}
	}
}
