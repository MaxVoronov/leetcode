package powx_n

import (
	"math"
	"testing"
)

func Test_myPow(t *testing.T) {
	tests := []struct {
		x    float64
		n    int
		want float64
	}{
		{x: 3., n: 3, want: 27.},
		{x: 2., n: 10, want: 1024.},
		{x: 2.1, n: 3, want: 9.261},
		{x: 2., n: -2, want: 0.25},
		{x: -2., n: 2, want: 4},
		{x: 123, n: 0, want: 1.},
		{x: 0, n: 0, want: 1.},
		{x: 0.00001, n: 2147483647, want: 0.},
		{x: 1., n: -2147483648, want: 1.},
		{x: -1., n: -2147483647, want: -1.},
		{x: -1., n: -2147483648, want: 1.},
		{x: 0.999999999, n: -2147483648, want: 8.56328},
	}

	for i, tc := range tests {
		result := myPow(tc.x, tc.n)
		if math.Abs(tc.want-result) > 1e-5 {
			t.Fatalf("Case #%d: expected: %v, got: %v", i+1, tc.want, result)
		}
	}
}
