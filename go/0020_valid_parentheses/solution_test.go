package valid_parentheses

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{input: "()", want: true},
		{input: "()[]{}", want: true},
		{input: "{[]}", want: true},
		{input: "(]", want: false},
		{input: "([)]", want: false},
		{input: ")(][}{", want: false},
		{input: "(()", want: false},
		{input: "}", want: false},
		{input: "", want: true},
	}

	for i, tc := range cases {
		got := isValid(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
