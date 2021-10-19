package binary_search

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 3, want: 2},
	}

	for i, tc := range cases {
		got := Search(tc.nums, tc.target)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
