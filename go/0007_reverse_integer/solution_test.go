package reverse_integer

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
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
	{input: -2_147_483_648, want: 0},
	{input: -123_456_789_876, want: 0},
}

var benchCases = []struct {
	input int
}{
	{input: 123},
	{input: -123},
	{input: 1_363_847_412},
	{input: 1_563_847_412},
}

func TestReverse(t *testing.T) {
	for i, tc := range testCases {
		got := Reverse(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}

func TestReverseV2(t *testing.T) {
	for i, tc := range testCases {
		got := ReverseV2(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}

func Benchmark_Reverse(b *testing.B) {
	for i, bc := range benchCases {
		b.Run(fmt.Sprintf("Case_%d", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Reverse(bc.input)
			}
		})
	}
}

func Benchmark_ReverseV2(b *testing.B) {
	for i, bc := range benchCases {
		b.Run(fmt.Sprintf("Case_%d", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ReverseV2(bc.input)
			}
		})
	}
}
