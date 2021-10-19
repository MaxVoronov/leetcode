package palindrome_number

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{input: 121, want: true},
		{input: -121, want: false},
		{input: 10, want: false},
		{input: 99, want: true},
		{input: -101, want: false},
		{input: 987789, want: true},
		{input: 9876789, want: true},
		{input: 1, want: true},
	}

	for i, tc := range tests {
		got := isPalindrome(tc.input)
		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
