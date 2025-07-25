// Package maths provides basic mathematical functions and acts a supplement to the math package.
package maths

import (
	"errors"
	"fmt"
	"math/big"
)

var ErrNLessThanK = errors.New("to calculate n choose k, n must be larger than or equal to k")

// Factorial returns the factorial of n, where n >= 0, with overflow detection.
// It returns an ErrOverflowDetected error if the calculation results in an overflow.
// If an overflow error is detected, the function returns 0, ErrOverflowDetected.
// In this case, consider *big.Int.MulRange() from the math/big package.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("n: %d. %w", n, ErrNegativeNumber)
	}

	if n > 20 {
		return 0, fmt.Errorf("the result of %d! is too large to hold in an int variable: %w", n, ErrOverflowDetected)
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result, nil
}

// Binomial returns the binomial coefficient of (n, k), n choose k, where n >= 0, k >= 0 and n >= k.
// It returns an ErrOverflowDetected error if the calculation results in an overflow.
// In this case, consider *big.Int.Binomial() from the math/big package.
func Binomial(n, k int) (int, error) {
	if n < 0 || k < 0 {
		return 0, fmt.Errorf("n: %d. k: %d. %w", n, k, ErrNegativeNumber)
	}

	if n < k {
		return 0, fmt.Errorf("n: %d. k: %d. %w", n, k, ErrNLessThanK)
	}

	result := new(big.Int).Binomial(int64(n), int64(k))

	if !result.IsInt64() {
		return 0, fmt.Errorf("the result of (%d, %d) is too large to hold in an int variable: %w", n, k, ErrOverflowDetected)
	}

	return int(result.Int64()), nil
}
