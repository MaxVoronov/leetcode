package climbing_stairs

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{input: 1, want: 1},
		{input: 2, want: 2},
		{input: 3, want: 3},
		{input: 5, want: 8},
		{input: 45, want: 1836311903},
	}

	for i, tc := range tests {
		got := climbStairs(tc.input)
		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
