package valid_anagram

import "testing"

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		inputS string
		inputT string
		want   bool
	}{
		{inputS: "anagram", inputT: "nagaram", want: true},
		{inputS: "restful", inputT: "fluster", want: true},
		{inputS: "rat", inputT: "car", want: false},
		{inputS: "x", inputT: "x", want: true},
		{inputS: "a", inputT: "ab", want: false},
		{inputS: "aa", inputT: "bb", want: false},
		{inputS: "aac", inputT: "bcb", want: false},
		{inputS: "ab", inputT: "ba", want: true},
	}

	for i, tc := range tests {
		got := isAnagram(tc.inputS, tc.inputT)
		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
