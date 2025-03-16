package maths

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
)

// Max returns the maximal value in a group of comparable values. It is a wrapper for the builtin slices.Max function.
func Max[T cmp.Ordered](numbers ...T) T {
	return slices.Max(numbers)
}

var ErrOverflowDetected = errors.New("arithmetic overflow detected")
var ErrAbsoluteValueOfMinInt = fmt.Errorf("can not calculate the absolute value of math.MinInt and store in an int variable: %w", ErrOverflowDetected)

// Abs returns |x|. Returns an error if calculating the absolute value of math.MinInt.
// In this case, use Abs() from the math/big package.
func Abs(x int) (int, error) {
	if x == math.MinInt {
		return 0, ErrAbsoluteValueOfMinInt
	}

	if x < 0 {
		return -x, nil
	}
	return x, nil
}
