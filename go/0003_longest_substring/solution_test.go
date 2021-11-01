package longest_substring

import (
	"reflect"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "abcabcbb", want: 3},
		{input: "bbbbb", want: 1},
		{input: "pwwkew", want: 3},
		{input: " ", want: 1},
		{input: "", want: 0},
		{input: "aab", want: 2},
		{input: "dvdf", want: 3},
		{input: "anviaj", want: 5},
	}

	for i, tc := range tests {
		got := lengthOfLongestSubstring(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
