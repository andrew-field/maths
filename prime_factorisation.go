package maths

import "math"

// PrimeFactor is designed to hold a prime number as 'value' and the index of the prime number as it appears in a complete prime factorisation product.
type PrimeFactor struct {
	Value, Index int
}

// PrimeFactorisation sends the prime factorisation of |x| on a channel, in order.
// If x is 0 or 1, PrimeFactorisation(x) returns a PrimeFactor with value x, index 1.
func PrimeFactorisation(x int) <-chan PrimeFactor {
	factorisationCh := make(chan PrimeFactor)

	go func() {
		// These special cases are handled inside the go function to avoid blocking the thread.
		if x == math.MinInt { // Special case when x is equal to math.MinInt. In this case, getting the absolute value would return an error, but the prime factorisation of |math.MinInt|, 2⁶³, is known.
			factorisationCh <- PrimeFactor{2, 63}
			close(factorisationCh)
			return
		}

		if x < 0 { // Because math.MinInt case is checked above, this can not panic with an error.
			x = -x
		}

		// Special case for 0 and 1.
		if x == 0 || x == 1 {
			factorisationCh <- PrimeFactor{x, 1}
			close(factorisationCh)
			return
		}

		primeCh := GetPrimeNumbersBelowAndIncluding(x)

		index := 0
		// For each prime, see if it is a factor and if so, how many times/the index with which it appears.
		for val := range primeCh {
			for ; x%val == 0; index++ {
				x /= val
			}

			if index != 0 {
				factorisationCh <- PrimeFactor{val, index}

				// If found all factors then finish.
				if x == 1 {
					go func() {
						for range primeCh { // Drain the primeCh in case there are values left.
						}
					}()
					close(factorisationCh)
					return
				}

				index = 0
			}
		}
	}()

	return factorisationCh
}
