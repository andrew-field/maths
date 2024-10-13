package maths

import "math"

// NumberOfDivisorsBruteForce returns the number of (positive) divisors of x. Uses a brute force method.
func NumberOfDivisorsBruteForce(x int) int {
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
		numberOfDivisors--
	}

	return numberOfDivisors
}

// SumOfDivisorsBruteForce returns the sum of all (positive) divisors of x. Uses a brute force method.
func SumOfDivisorsBruteForce(x int) int {
	sumOfDivisors := 0
	for divisor := range GetDivisorsBruteForce(x) {
		sumOfDivisors += divisor
	}

	return sumOfDivisors
}

// GetDivisorsBruteForce fills a channel with all the (positive) divisors of x, unsorted. Uses a brute force method.
func GetDivisorsBruteForce(x int) <-chan int {
	divisorCh := make(chan int)

	go func() {
		if x == 0 {
			close(divisorCh)
			return
		}

		x = Abs(x)

		limit := int(math.Sqrt(float64(x)))

		for i := 1; i < limit; i++ {
			if x%i == 0 {
				divisorCh <- i
				divisorCh <- x / i
			}
		}

		if x%limit == 0 {
			divisorCh <- limit
			if limit*limit != x {
				divisorCh <- x / limit
			}
		}

		close(divisorCh)
	}()

	return divisorCh
}
