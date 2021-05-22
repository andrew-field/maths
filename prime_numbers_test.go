package maths

import (
	"testing"
)

// TestGetPrimeNumbers checks the first few prime numbers.
func TestGetPrimeNumbers(t *testing.T) {
	expectedPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	primeChannel := GetPrimeNumbers()

	for _, expectedPrime := range expectedPrimes {
		if actualPrime := <-primeChannel; actualPrime != expectedPrime {
			t.Errorf("Actual prime: %v. Expected prime: %v.", actualPrime, expectedPrime)
		}
	}
}
