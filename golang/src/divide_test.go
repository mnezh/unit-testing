package divide

import (
	"errors"
	"testing"
)

// Data-driven test: cases defined as data structure
func TestDivide(t *testing.T) {
	for _, tc := range []struct {
		testCaseName string
		dividend     float64
		divisor      float64
		want         float64
		wantErr      error
	}{
		{
			testCaseName: "divide integer",
			dividend:     10,
			divisor:      2,
			want:         5,
		},
		{
			testCaseName: "divide float",
			dividend:     10,
			divisor:      4,
			want:         2.5,
		},
		{
			testCaseName: "divide by zero",
			dividend:     10,
			divisor:      0,
			wantErr:      DivisionByZero,
		},
	} {
		t.Run(tc.testCaseName, func(subt *testing.T) {
			res, err := Divide(tc.dividend, tc.divisor)
			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					subt.Errorf("Divide(%v,%v): unexpected error '%v' instead of '%v'", tc.dividend, tc.divisor, err, tc.wantErr)
				}
			} else if res != tc.want {
				subt.Errorf("Divide(%v,%v): %v != %v", tc.dividend, tc.divisor, res, tc.want)
			}
		})
	}
}
