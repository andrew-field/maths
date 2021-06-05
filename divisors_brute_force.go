package maths

import "math"

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
		numberOfDivisors--
	}

	return numberOfDivisors
}

// SumOfDivisors2 returns the sum of all (positive) divisors of x. Uses a brute force method.
func SumOfDivisors2(x int) int {
	sumOfDivisors := 0
	for divisor := range Divisors2(x) {
		sumOfDivisors += divisor
	}

	return sumOfDivisors
}

// Divisors2 fills a channel with all the (positive) divisors of x, unsorted. Uses a brute force method.
func Divisors2(x int) <-chan int {
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
