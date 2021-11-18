package sqrtx

import "testing"

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		input int32
		want  int
	}{
		{input: 2147395599, want: 46339},
		{ input: 16, want: 4 },
		{ input: 9, want: 3 },
		{ input: 8, want: 2 },
		{ input: 7, want: 2 },
		{ input: 4, want: 2 },
		{ input: 1, want: 1 },
		{ input: 0, want: 0 },
	}

	for i, tc := range tests {
		got := mySqrt(int(tc.input))
		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
