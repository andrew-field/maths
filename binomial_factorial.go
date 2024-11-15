package maths

import (
	"fmt"
	"math"
)

// Factorial returns the factorial of |n|, with overflow detection.
// If an overflow error is detected when the numbers get too large, the function returns 0, ErrOverflowDetected.
// In this case, use *bigInt.MulRange() from the math/big package.
func Factorial(n int) (int, error) {
	absN, err := Abs(n)
	if err != nil {
		return 0, fmt.Errorf("failed to get Abs(%d): %w", n, err)
	}

	result, err := fact(absN)
	if err != nil {
		return 0, fmt.Errorf("failed to get fact(%d): %w", absN, err)
	}
	return result, nil
}

func fact(n int) (int, error) {
	if n == 0 {
		return 1, nil
	}

	// Recursive call for n-1 factorial.
	partial, err := fact(n - 1)
	if err != nil {
		return 0, err // propagate overflow error.
	}

	// Check for overflow before multiplication. Partial can only be positive here.
	if n > math.MaxInt/partial {
		return 0, fmt.Errorf("failed to calculate %d * %d. The result is too large to hold in an int variable: %w", n, partial, ErrOverflowDetected)
	}

	return n * partial, nil
}

// Binomial returns the binomial coefficient of (|n|, |k|), |n| choose |k|, where |n| >= |k|.
// If an overflow error is detected when the numbers get too large, the function returns 0, ErrOverflowDetected.
// In this case, use *bigInt.Binomial() from the math/big package.
func Binomial(n, k int) (int, error) {
	absN, err := Abs(n)
	if err != nil {
		return 0, fmt.Errorf("failed to get Abs(%d): %w", n, err)
	}
	absK, err := Abs(k)
	if err != nil {
		return 0, fmt.Errorf("failed to get Abs(%d): %w", k, err)
	}

	// |n| must be larger than or equal to |k|.
	differenceOfAbsolutes := absN - absK
	if differenceOfAbsolutes < 0 {
		return 0, fmt.Errorf("to calculate |n| choose |k|, |n| must be larger than or equal to |k|. |n|:%d. |k|:%d", n, k)
	}

	// Calculate the factorial of |n|.
	factN, err := fact(absN)
	if err != nil {
		return 0, fmt.Errorf("failed to get fact(%d): %w", absN, err)
	}

	// Calculate the factorial of |k|.
	factK, err := fact(absK)
	if err != nil {
		return 0, fmt.Errorf("failed to get fact(%d): %w", absK, err)
	}

	// Calculate the factorial of |n| - |k|.
	factDifference, err := fact(differenceOfAbsolutes)
	if err != nil {
		return 0, fmt.Errorf("failed to get fact(%d): %w", differenceOfAbsolutes, err)
	}

	return factN / factK / factDifference, nil // If the preceding steps did not give an error, this calculation can not give an error.
}
