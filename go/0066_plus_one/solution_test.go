package plus_one

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{4, 3, 9, 9}, want: []int{4, 4, 0, 0}},
		{input: []int{1, 2, 3}, want: []int{1, 2, 4}},
		{input: []int{4, 3, 2, 1}, want: []int{4, 3, 2, 2}},
		{input: []int{9}, want: []int{1, 0}},
		{input: []int{9, 9, 9, 9, 9, 9, 9, 9, 9}, want: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{input: []int{0}, want: []int{1}},
	}

	for i, tc := range tests {
		got := plusOne(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
