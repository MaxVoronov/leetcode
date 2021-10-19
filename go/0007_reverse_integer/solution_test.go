package reverse_integer

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{input: 123, want: 321},
		{input: -123, want: -321},
		{input: 120, want: 21},
		{input: 5, want: 5},
		{input: -7, want: -7},
		{input: 0, want: 0},
		// Int32 overflow
		{input: 1_363_847_412, want: 2_147_483_631},
		{input: 1_563_847_412, want: 0},
		{input: 1_534_236_469, want: 0},
		{input: 123_456_789_876, want: 0},
	}

	for i, tc := range tests {
		got := Reverse(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
