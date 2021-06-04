package maths

import (
	"testing"
)

// TestGetPrimeNumbers checks the first few prime numbers.
func TestGetPrimeNumbers(t *testing.T) {
	expectedPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	primeChannel, doneCh := GetPrimeNumbers()

	for _, expectedPrime := range expectedPrimes {
		if actualPrime := <-primeChannel; actualPrime != expectedPrime {
			t.Errorf("Actual prime: %v. Expected prime: %v.", actualPrime, expectedPrime)
		}
	}

	doneCh <- true
}

func TestGetPrimeNumbersBelow(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult []int
	}{
		{-10, []int{2, 3, 5, 7}},
		{-3, []int{2}},
		{-2, []int{}},
		{-1, []int{}},
		{0, []int{}},
		{1, []int{}},
		{2, []int{}},
		{3, []int{2}},
		{10, []int{2, 3, 5, 7}},
		{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}

	for _, tC := range testCases {
		primeChannel := GetPrimeNumbersBelow(tC.input)

		for _, expectedPrime := range tC.expectedResult {
			if actualPrime := <-primeChannel; actualPrime != expectedPrime {
				t.Errorf("Input in test: %v. Actual prime: %v. Expected prime: %v.", tC.input, actualPrime, expectedPrime)
			}
		}

		// Check the prime channel does not have too many values.
		if prime, more := <-primeChannel; more {
			t.Errorf("Received more primes than expected. Input in test: %v. Unexpected prime: %v", tC.input, prime)
		}
	}
}
