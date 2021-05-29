package maths

// NumberOfDivisors returns the number of (positive) divisors of x. Uses PrimeFactorisation(x).
// Does not handle math.MinInt64.
func NumberOfDivisors(x int) int {
	x = Abs(x)

	if x == 0 || x == 1 {
		return x
	}

	factorisationChannel := PrimeFactorisation(x)

	// Calculate the number of divisors.
	divisors := 1
	for primeFactor := range factorisationChannel {
		divisors *= primeFactor.index + 1
	}

	return divisors
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
			for i := 1; i <= primeFactor.index; i++ {
				for sectionLimit := sectionLength * i; j < sectionLimit; j++ {
					existingDivisors = append(existingDivisors, existingDivisors[j]*primeFactor.value)
					divisorCh <- existingDivisors[sectionLength+j]
				}
			}
		}

		close(divisorCh)
	}()

	return divisorCh
}
