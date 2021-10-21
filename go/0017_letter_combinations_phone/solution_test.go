package letter_combinations_phone

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	cases := []struct {
		input string
		want  []string
	}{
		{input: "23", want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{input: "2", want: []string{"a", "b", "c"}},
		{input: "", want: []string{}},
	}

	for i, tc := range cases {
		got := letterCombinations(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
