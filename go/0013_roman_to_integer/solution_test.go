package roman_to_integer

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "I", want: 1},
		{input: "III", want: 3},
		{input: "IV", want: 4},
		{input: "VI", want: 6},
		{input: "IX", want: 9},
		{input: "LVIII", want: 58},
		{input: "MCMXCIV", want: 1994},
		{input: "MMMCMXCIX", want: 3999},
		{input: "MMMCCCXXXIII", want: 3333},
	}

	for i, tc := range tests {
		got := romanToInt(tc.input)
		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
