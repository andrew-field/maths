package maths

import (
	"fmt"
	"math"
)

// NumberOfDivisors returns the number of (positive) divisors of x. Uses PrimeFactorisation().
func NumberOfDivisors(x int) int {
	if x == math.MinInt { // Special case when x is equal to math.MinInt. In this case, getting the absolute value would return an error, but the number of divisors of |math.MinInt|, 2⁶³, is known.
		return 64
	}

	if x < 0 { // Because math.MinInt case is checked above, this can not panic with an error.
		x = -x
	}

	// Special cases for 0 and 1.
	if x == 0 || x == 1 {
		return x
	}

	// Calculate the number of divisors.
	numberOfDivisors := 1
	for primeFactor := range PrimeFactorisation(x) {
		numberOfDivisors *= primeFactor.Index + 1
	}

	return numberOfDivisors
}

// GetDivisors fills a channel with all the (positive) divisors of x, unsorted. Uses PrimeFactorisation().
func GetDivisors(x int) (<-chan int, error) {
	divisorCh := make(chan int)

	x, err := Abs(x) // A special case can not be made for x equal to math.MinInt because |math.MinInt| itself would be a positive divisor of math.MinInt, which can not be stored in an int variable.
	if err != nil {
		close(divisorCh)
		return divisorCh, fmt.Errorf("failed to get Abs(%d): %w", x, err)
	}

	if x == 0 {
		close(divisorCh)
		return divisorCh, nil
	}

	go func() {
		divisorCh <- 1

		if x == 1 {
			close(divisorCh)
			return
		}

		existingDivisors := []int{1}

		for primeFactor := range PrimeFactorisation(x) {
			sectionLength := len(existingDivisors)

			// For each new prime factor found, multiply all of the existing divisors by the prime factor and add them to the slice of existing divisors.
			// Repeat this step, using the updated slice of existing divisors, as many times as the prime factor index of that prime factor, as it appears in the prime factorisation of x.
			// In this way, all of the possible product combinations of the prime factors and there respective indexes, are calculated.
			j := 0
			for i := 1; i <= primeFactor.Index; i++ {
				for sectionLimit := sectionLength * i; j < sectionLimit; j++ {
					existingDivisors = append(existingDivisors, existingDivisors[j]*primeFactor.Value)
					divisorCh <- existingDivisors[sectionLength+j]
				}
			}
		}

		close(divisorCh)
	}()

	return divisorCh, nil
}

// SumOfDivisors returns the sum of all the (positive) divisors of x. Uses PrimeFactorisation().
// It returns an error wrapping ErrOverflowDetected if the calculation results in an overflow.
// In this case, you can try SumOfDivisorsBruteForce(), although this might still give an error.
func SumOfDivisors(x int) (int, error) {
	x, err := Abs(x) // A special case can not be made for x equal to math.MinInt because the sum of all positive divisors of math.MinInt can not be stored in an int variable.
	if err != nil {
		return 0, fmt.Errorf("failed to get Abs(%d): %w", x, err)
	}

	if x == 0 || x == 1 {
		return x, nil
	}

	// Calculate the sum of divisors.
	sumOfDivisors := 1
	for primeFactor := range PrimeFactorisation(x) {
		primeFactorPow, err := Pow(primeFactor.Value, primeFactor.Index+1)
		if err != nil {
			return 0, fmt.Errorf("failed to get Pow(%d, %d): %w", primeFactor.Value, primeFactor.Index+1, err)
		}

		// Check for overflow before multiplication. Partial can only be positive here.
		partial := (primeFactorPow - 1) / (primeFactor.Value - 1)
		if sumOfDivisors > math.MaxInt/partial {
			return 0, fmt.Errorf("failed to calculate %d * %d. The result is too large to hold in an int variable: %w", sumOfDivisors, partial, ErrOverflowDetected)
		}
		sumOfDivisors *= partial
	}

	return sumOfDivisors, nil
}
