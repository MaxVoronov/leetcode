package add_two_numbers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{l1: []int{2, 4, 3}, l2: []int{5, 6, 4}, want: []int{7, 0, 8}},
		{l1: []int{0}, l2: []int{0}, want: []int{0}},
		{l1: []int{9, 9, 9, 9, 9, 9, 9}, l2: []int{9, 9, 9, 9}, want: []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for i, tc := range tests {
		got := AddTwoNumbers(SliceToLinkedList(tc.l1), SliceToLinkedList(tc.l2))
		fmt.Printf("List1: %+v\n", tc.l1)
		fmt.Printf("List2: %+v\n", tc.l2)
		fmt.Printf("Result: %+v\n", got.ToSlice())

		if !reflect.DeepEqual(tc.want, got.ToSlice()) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got.ToSlice())
		}
	}
}
