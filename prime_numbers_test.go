package maths

import (
	"fmt"
	"testing"
)

// TestGetPrimeNumbers checks the first few prime numbers.
func TestGetPrimeNumbers(t *testing.T) {
	expectedPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	primeChannel, doneCh := GetPrimeNumbers()

	for _, expectedPrime := range expectedPrimes {
		if actualPrime := <-primeChannel; actualPrime != expectedPrime {
			t.Errorf("Actual prime: %d. Expected prime: %d.", actualPrime, expectedPrime)
		}
	}

	doneCh <- true
}

func TestGetPrimeNumbersBelowAndIncluding(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult []int
	}{
		{-10, []int{2, 3, 5, 7}},
		{-3, []int{2, 3}},
		{-2, []int{2}},
		{-1, []int{}},
		{0, []int{}},
		{1, []int{}},
		{2, []int{2}},
		{3, []int{2, 3}},
		{4, []int{2, 3}},
		{5, []int{2, 3, 5}},
		{6, []int{2, 3, 5}},
		{7, []int{2, 3, 5, 7}},
		{8, []int{2, 3, 5, 7}},
		{9, []int{2, 3, 5, 7}},
		{10, []int{2, 3, 5, 7}},
		{11, []int{2, 3, 5, 7, 11}},
		{12, []int{2, 3, 5, 7, 11}},
		{13, []int{2, 3, 5, 7, 11, 13}},
		{14, []int{2, 3, 5, 7, 11, 13}},
		{15, []int{2, 3, 5, 7, 11, 13}},
		{16, []int{2, 3, 5, 7, 11, 13}},
		{17, []int{2, 3, 5, 7, 11, 13, 17}},
		{18, []int{2, 3, 5, 7, 11, 13, 17}},
		{19, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{21, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{22, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{23, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{24, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{25, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{26, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{27, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{28, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{29, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{31, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}},
		{32, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}},
		{33, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}},
		{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
		{101, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			primeCh := GetPrimeNumbersBelowAndIncluding(tC.input)

			for _, expectedPrime := range tC.expectedResult {
				if actualPrime := <-primeCh; actualPrime != expectedPrime {
					t.Errorf("Actual prime: %d. Expected prime: %d.", actualPrime, expectedPrime)
				}
			}

			// Check the prime channel does not have too many values.
			if prime, more := <-primeCh; more {
				t.Errorf("Received more primes than expected. Unexpected prime: %d", prime)
			}
		})
	}
}
