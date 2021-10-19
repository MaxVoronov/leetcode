package reverse_string

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		input []byte
		want  []byte
	}{
		{input: []byte{'h', 'e', 'l', 'l', 'o'}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
		{input: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, want: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		{input: []byte{}, want: []byte{}},
		{input: []byte{'A'}, want: []byte{'A'}},
	}

	for i, tc := range cases {
		result := append([]byte{}, tc.input...)
		reverseString(result)
		if !reflect.DeepEqual(tc.want, result) {
			t.Fatalf("Case #%d: expected: %s, got: %s", i+1, string(tc.want), string(result))
		}
	}
}
