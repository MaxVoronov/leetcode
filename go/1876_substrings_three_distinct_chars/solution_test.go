package substrings_three_distinct_chars

import (
	"reflect"
	"testing"
)

func Test_countGoodSubstrings(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "xyzzaz", want: 1},
		{input: "aababcabc", want: 4},
	}

	for i, tc := range tests {
		got := countGoodSubstrings(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
