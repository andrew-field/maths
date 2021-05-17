package maths

type PrimeFactor struct {
	value int
	index int
}

// PrimeFactorisation sends the prime factorisation of a number on a channel, in ascending order.
// If x is negative, PrimeFactorisation(x) returns PrimeFactorisation(-x)
// If x is 0 or 1, PrimeFactorisation(x) returns factor x, index 1.
// Does not handle math.MinInt64.
func PrimeFactorisation(number int) <-chan PrimeFactor {
	factorisationCh := make(chan PrimeFactor)

	go func() {
		if number < 0 {
			number = -number
		}

		// Special case for 0 and 1.
		if number == 0 || number == 1 {
			factorisationCh <- PrimeFactor{number, 1}
			close(factorisationCh)
			return
		}

		primeChannel := GetPrimeNumbers()

		// For each prime, see if it is a factor.
		for val := range primeChannel {
			index := 0
			for ; number%val == 0; index++ {
				number /= val
			}

			if index != 0 {
				factorisationCh <- PrimeFactor{val, index}

				// If found all factors then finish.
				if number == 1 {
					close(factorisationCh)
					return
				}
			}
		}
	}()

	return factorisationCh
}
