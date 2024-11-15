package maths

import (
	"fmt"
	"math"
)

// NumberOfDivisorsBruteForce returns the number of (positive) divisors of x. Uses a brute force method.
func NumberOfDivisorsBruteForce(x int) int {
	if x == 0 {
		return 0
	}

	if x == math.MinInt { // Special case when x is equal to math.MinInt. In this case, getting the absolute value would return an error, but the number of divisors of |math.MinInt|, 2⁶³, is known.
		return 64
	}

	if x < 0 { // Because math.MinInt case is checked above, this can not panic with an error.
		x = -x
	}

	// Every divisor of x has a corresponding divisor of x, which when multiplied together will equal x.
	// Each pair of divisors will lie either side of the SQRT of x (possibly equal to SQRT of x for square numbers).
	// Therefore, one only needs to check up to the SQRT of x to know the number of divisors.
	limit := int(math.Sqrt(float64(x))) // The int conversion returns the floor of the square root (as an int).

	numberOfDivisors := 1
	for i := 2; i <= limit; i++ {
		if x%i == 0 {
			numberOfDivisors++
		}
	}

	numberOfDivisors *= 2

	if limit*limit == x { // If a square number.
		numberOfDivisors-- // Don't count the same divisor twice.
	}

	return numberOfDivisors
}

// GetDivisorsBruteForce fills a channel with all the (positive) divisors of x, unsorted. Uses a brute force method.
func GetDivisorsBruteForce(x int) (<-chan int, error) {
	divisorCh := make(chan int)

	if x == 0 {
		close(divisorCh)
		return divisorCh, nil
	}

	// A special case can not be made for x equal to math.MinInt because |math.MinInt| itself would be a positive divisor of math.MinInt, which can not be stored in an int variable.
	x, err := Abs(x)
	if err != nil {
		close(divisorCh)
		return divisorCh, fmt.Errorf("failed to get Abs(%d): %w", x, err)
	}

	go func() {
		// Every divisor of x has a corresponding divisor of x, which when multiplied together will equal x.
		// Each pair of divisors will lie either side of the SQRT of x (possibly equal to SQRT of x for square numbers).
		// Therefore, one only needs to check up to the SQRT of x to be able to find all of the divisors.
		limit := int(math.Sqrt(float64(x))) // The int conversion returns the floor of the square root (as an int).

		for i := 1; i < limit; i++ {
			if x%i == 0 {
				divisorCh <- i
				divisorCh <- x / i
			}
		}

		if x%limit == 0 {
			divisorCh <- limit
			if limit*limit != x { // If a square number, include the SQRT of x only once.
				divisorCh <- x / limit
			}
		}

		close(divisorCh)
	}()

	return divisorCh, nil
}

// SumOfDivisorsBruteForce returns the sum of all (positive) divisors of x. Uses GetDivisorsBruteForce().
// If an overflow error is detected when the numbers get too large, the function returns 0, ErrOverflowDetected.
func SumOfDivisorsBruteForce(x int) (int, error) {
	divisorCh, err := GetDivisorsBruteForce(x)
	if err != nil {
		return 0, fmt.Errorf("failed to get GetDivisorsBruteForce(%d): %w", x, err)
	}

	sumOfDivisors := 0
	for divisor := range divisorCh {
		// Check for overflow before addition.
		if sumOfDivisors > math.MaxInt-divisor {
			return 0, fmt.Errorf("failed to calculate %d + %d. The result is too large to hold in an int variable: %w", sumOfDivisors, divisor, ErrOverflowDetected)
		}
		sumOfDivisors += divisor
	}

	return sumOfDivisors, nil
}
