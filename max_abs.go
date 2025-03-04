package maths

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

// Max returns the maximum int from a group of integers.
// Max() = 0.
func Max(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	return slices.Max(numbers)
}

// ErrOverflowDetected is an error that indicates an arithmetic overflow has been detected.
var ErrOverflowDetected = errors.New("arithmetic overflow detected")

// ErrAbsoluteValueOfMinInt is an error that indicates an attempt to calculate
// the absolute value of math.MinInt and store it in an int variable, which is
// not possible due to overflow.
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
