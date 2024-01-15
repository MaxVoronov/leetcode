package merge_sorted_array

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type input struct {
		nums1, nums2 []int
		m, n         int
	}

	tests := []struct {
		input input
		want  []int
	}{
		{
			input: input{
				nums1: []int{1, 2, 3, 0, 0, 0},
				nums2: []int{2, 5, 6},
				m:     3,
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			input: input{
				nums1: []int{1},
				nums2: []int{},
				m:     1,
				n:     0,
			},
			want: []int{1},
		},
		{
			input: input{
				nums1: []int{0},
				nums2: []int{1},
				m:     0,
				n:     1,
			},
			want: []int{1},
		},
		{
			input: input{
				nums1: []int{7, 8, 9, 0, 0},
				nums2: []int{2, 3},
				m:     3,
				n:     2,
			},
			want: []int{2, 3, 7, 8, 9},
		},
	}

	for i, tc := range tests {
		merge(tc.input.nums1, tc.input.m, tc.input.nums2, tc.input.n)
		if !reflect.DeepEqual(tc.want, tc.input.nums1) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, tc.input.nums1)
		}
	}
}
