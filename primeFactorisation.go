package maths

type PrimeFactor struct {
	value int
	index int
}

// PrimeFactorisation sends the prime factorisation of a number on a channel, in ascending order.
// If x is negative, PrimeFactorisation(x) returns PrimeFactorisation(-x)
// If x is 0 or 1, PrimeFactorisation(x) returns factor x, index 1.
func PrimeFactorisation(numberToFactorise int) <-chan PrimeFactor {
	factorisationChannel := make(chan PrimeFactor)

	go func() {
		if numberToFactorise < 0 {
			numberToFactorise = -numberToFactorise
		}

		// Special case for 0 and 1.
		if numberToFactorise == 0 || numberToFactorise == 1 {
			factorisationChannel <- PrimeFactor{numberToFactorise, 1}
			close(factorisationChannel)
			return
		}

		primeChannel := GetPrimeNumbers()

		// For each prime, see if it is a factor.
		for val := range primeChannel {
			index := 0
			for ; numberToFactorise%val == 0; index++ {
				numberToFactorise /= val
			}

			if index != 0 {
				factorisationChannel <- PrimeFactor{val, index}

				// If found all factors then finish.
				if numberToFactorise == 1 {
					close(factorisationChannel)
					return
				}
			}
		}
	}()

	return factorisationChannel
}
