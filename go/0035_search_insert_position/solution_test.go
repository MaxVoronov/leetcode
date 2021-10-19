package search_insert_position

import (
	"reflect"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		want   int
	}{
		{input: []int{1, 3, 5, 6, 8}, target: 7, want: 4},
		{input: []int{1, 3, 5, 6, 8}, target: 5, want: 2},
		{input: []int{-12, -8, -2, 0, 3, 12, 33, 48, 59, 64, 72}, target: -5, want: 2},
		{input: []int{1, 3, 5, 6}, target: 5, want: 2},
		{input: []int{1, 3, 5, 6}, target: 2, want: 1},
		{input: []int{1, 3, 5, 6}, target: 7, want: 4},
		{input: []int{1, 3, 5, 6}, target: 0, want: 0},
		{input: []int{1}, target: 0, want: 0},
		{input: []int{}, target: 0, want: 0},
	}

	for i, tc := range cases {
		got := searchInsert(tc.input, tc.target)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
