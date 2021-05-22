package maths

// PrimeFactor is designed to hold a prime number as 'value' and the index of the prime number when it appears in a complete prime factorisation product.
type PrimeFactor struct {
	value, index int
}

// PrimeFactorisation(x) sends the prime factorisation of |x| on a channel, in ascending order.
// If x is 0 or 1, PrimeFactorisation(x) returns a PrimeFactor with value x, index 1.
// Does not handle math.MinInt64.
func PrimeFactorisation(number int) <-chan PrimeFactor {
	factorisationCh := make(chan PrimeFactor)

	go func() {
		number = Abs(number)

		// Special case for 0 and 1.
		if number == 0 || number == 1 {
			factorisationCh <- PrimeFactor{number, 1}
			close(factorisationCh)
			return
		}

		primeChannel := GetPrimeNumbers()

		index := 0
		// For each prime, see if it is a factor and if so, how many times/the index with which it appears.
		for val := range primeChannel {
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

				index = 0
			}
		}
	}()

	return factorisationCh
}
