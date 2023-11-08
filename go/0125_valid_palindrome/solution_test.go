package valid_palindrome

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	input string
	want  bool
}{
	{input: "A man, a plan, a canal: Panama", want: true},
	{input: "race a car", want: false},
	{input: "0P", want: false},
	{input: ".a", want: true},
	{input: " ", want: true},
}

func Benchmark_isPalindrome(b *testing.B) {
	for i, tc := range testCases {
		b.Run(fmt.Sprintf("Case_%d", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPalindrome(tc.input)
			}
		})
	}
}

func Benchmark_isPalindromeV2(b *testing.B) {
	for i, tc := range testCases {
		b.Run(fmt.Sprintf("Case_%d", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPalindromeV2(tc.input)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	for i, tc := range testCases {
		got := isPalindrome(tc.input)
		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}

func Test_isPalindromeV2(t *testing.T) {
	for i, tc := range testCases {
		got := isPalindromeV2(tc.input)
		if tc.want != got {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
