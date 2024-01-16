package remove_duplicates_from_array

import (
	"reflect"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		input   []int
		want    int
		wantArr []int
	}{
		{input: []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5, wantArr: []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0}},
		{input: []int{1, 1, 2}, want: 2, wantArr: []int{1, 2, 0}},
		{input: []int{1, 2}, want: 2, wantArr: []int{1, 2}},
		{input: []int{1}, want: 1, wantArr: []int{1}},
		{input: []int{}, want: 0, wantArr: []int{}},
	}

	for i, tc := range tests {
		got := removeDuplicates(tc.input)

		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}

		if !reflect.DeepEqual(tc.wantArr[:got], tc.input[:got]) {
			t.Fatalf("Case #%d: invalid modified slice - expected: %v, got: %v", i+1, tc.wantArr[:got], tc.input[:got])
		}
	}
}
