package maths

// NumberOfDivisors returns the number of (positive) divisors.
func NumberOfDivisors(x int) int {
	if x < 0 {
		x = -x
	}

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

// Divisors fills a channel with all the (positive) divisors of a number.
func Divisors(number int) <-chan int {
	divisorChannel := make(chan int)

	go func() {
		if number == 0 {
			close(divisorChannel)
			return
		}

		divisorChannel <- 1

		if number < 0 {
			number = -number
		}

		if number == 1 {
			close(divisorChannel)
			return
		}

		existingDivisors := []int{1}

		factorisationChannel := PrimeFactorisation(number)

		for primeFactor := range factorisationChannel {
			sectionLength := len(existingDivisors)

			j := 0
			for i := 1; i <= primeFactor.index; i++ {
				for sectionLimit := sectionLength * i; j < sectionLimit; j++ {
					existingDivisors = append(existingDivisors, existingDivisors[j]*primeFactor.value)
					divisorChannel <- existingDivisors[sectionLimit+j]
				}
			}
		}

		close(divisorChannel)
	}()

	return divisorChannel
}