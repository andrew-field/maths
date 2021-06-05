package maths

import (
	"math"
)

// NumberOfDivisors returns the number of (positive) divisors of x. Uses PrimeFactorisation(x).
// Does not handle math.MinInt64.
func NumberOfDivisors(x int) int {
	x = Abs(x)

	if x == 0 || x == 1 {
		return x
	}

	factorisationChannel := PrimeFactorisation(x)

	// Calculate the number of divisors.
	numberOfDivisors := 1
	for primeFactor := range factorisationChannel {
		numberOfDivisors *= primeFactor.Index + 1
	}

	return numberOfDivisors
}

// NumberOfDivisors2 returns the number of (positive) divisors of x. Uses a brute force method.
func NumberOfDivisors2(x int) int {
	if x == 0 {
		return 0
	}
	x = Abs(x)

	limit := int(math.Sqrt(float64(x)))

	numberOfDivisors := 0
	for i := 1; i <= limit; i++ {
		if x%i == 0 {
			numberOfDivisors++
		}
	}

	numberOfDivisors *= 2

	if limit*limit == x {
		numberOfDivisors -= 1
	}

	return numberOfDivisors
}

// SumOfDivisors returns the sum of all (positive) divisors of x. Uses PrimeFactorisation(x).
// Does not handle math.MinInt64.
func SumOfDivisors(x int) int {
	x = Abs(x)

	if x == 0 || x == 1 {
		return x
	}

	factorisationChannel := PrimeFactorisation(x)

	// Calculate the number of divisors.
	sumOfDivisors := 1
	for primeFactor := range factorisationChannel {
		sumOfDivisors *= (Pow(primeFactor.Value, primeFactor.Index+1) - 1) / (primeFactor.Value - 1)
	}

	return sumOfDivisors
}

// Divisors fills a channel with all the (positive) divisors of x. Uses PrimeFactorisation(x).
// Does not handle math.MinInt64.
func Divisors(x int) <-chan int {
	divisorCh := make(chan int)

	go func() {
		if x == 0 {
			close(divisorCh)
			return
		}

		divisorCh <- 1

		x = Abs(x)

		if x == 1 {
			close(divisorCh)
			return
		}

		existingDivisors := []int{1}

		factorisationChannel := PrimeFactorisation(x)

		for primeFactor := range factorisationChannel {
			sectionLength := len(existingDivisors)

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

	return divisorCh
}
