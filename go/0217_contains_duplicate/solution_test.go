package contains_duplicate

import "testing"

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		input []int
		want  bool
	}{
		{input: []int{1, 2, 3, 1}, want: true},
		{input: []int{1, 2, 3, 4}, want: false},
		{input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	}

	for i, tc := range tests {
		got := containsDuplicate(tc.input)
		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
