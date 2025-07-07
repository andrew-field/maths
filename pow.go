package maths

import (
	"fmt"
	"math"
)

// Pow returns the x^|y|.
// Returns 1, nil for all x and y when y is 0 or x is 1.
// It returns an error wrapping ErrOverflowDetected if the calculation results in an overflow.
// In this case, consider *big.Int.Exp() from the math/big package.
func Pow(x, y int) (int, error) {
	if y == 0 || x == 1 {
		return 1, nil
	}

	absY, err := Abs(y)
	if err != nil {
		return 0, fmt.Errorf("failed to get Abs(%d): %w", y, err)
	}

	result := x
	for i := 1; i < absY; i++ {
		// Check for overflow before multiplication. x could be positive or negative.
		if result > math.MaxInt/x && result > math.MinInt/x {
			return 0, fmt.Errorf("failed to calculate %d * %d. The result is too large to hold in an int variable: %w", result, x, ErrOverflowDetected)
		}
		result *= x
	}

	return result, nil
}
