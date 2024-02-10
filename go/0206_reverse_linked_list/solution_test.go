package reverse_linked_list

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{input: []int{1, 2}, want: []int{2, 1}},
		{input: []int{}, want: []int{}},
		{input: []int{9, 9, 9, 9, 9}, want: []int{9, 9, 9, 9, 9}},
	}

	for i, tc := range tests {
		got := reverseList(SliceToLinkedList(tc.input))
		if !reflect.DeepEqual(tc.want, got.ToSlice()) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got.ToSlice())
		}
	}
}
