package divide

import (
	"errors"
	"fmt"
	"math"
)

var DivisionByZero = errors.New("Division by zero")

// Returns error instead of Infinity in case of division by zero
func Divide(a, b float64) (float64, error) {
	res := a / b
	if math.IsInf(res, 0) {
		return 0, fmt.Errorf("%v/%v: %w", a, b, DivisionByZero)

	}
	return res, nil
}
