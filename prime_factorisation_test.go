package maths

import (
	"fmt"
	"math"
	"testing"
)

func TestPrimeFactorisation(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult []PrimeFactor
	}{
		{math.MinInt, []PrimeFactor{{2, 63}}},
		{-5, []PrimeFactor{{5, 1}}},
		{-4, []PrimeFactor{{2, 2}}},
		{-3, []PrimeFactor{{3, 1}}},
		{-2, []PrimeFactor{{2, 1}}},
		{-1, []PrimeFactor{{1, 1}}},
		{0, []PrimeFactor{{0, 1}}},
		{1, []PrimeFactor{{1, 1}}},
		{2, []PrimeFactor{{2, 1}}},
		{3, []PrimeFactor{{3, 1}}},
		{4, []PrimeFactor{{2, 2}}},
		{5, []PrimeFactor{{5, 1}}},
		{6, []PrimeFactor{{2, 1}, {3, 1}}},
		{7, []PrimeFactor{{7, 1}}},
		{8, []PrimeFactor{{2, 3}}},
		{9, []PrimeFactor{{3, 2}}},
		{10, []PrimeFactor{{2, 1}, {5, 1}}},
		{100, []PrimeFactor{{2, 2}, {5, 2}}},
		{101, []PrimeFactor{{101, 1}}},
		{1000, []PrimeFactor{{2, 3}, {5, 3}}},
		{4561356, []PrimeFactor{{2, 2}, {3, 1}, {593, 1}, {641, 1}}},
		{600851475143, []PrimeFactor{{71, 1}, {839, 1}, {1471, 1}, {6857, 1}}},
		{math.MaxInt, []PrimeFactor{{7, 2}, {73, 1}, {127, 1}, {337, 1}, {92737, 1}, {649657, 1}}},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			primeFactorisationCh := PrimeFactorisation(tC.input)

			for _, expectedPrimeFactor := range tC.expectedResult {
				if actualPrimeFactor := <-primeFactorisationCh; actualPrimeFactor != expectedPrimeFactor {
					t.Errorf("Actual factor: %v. Expected factor: %v.", actualPrimeFactor, expectedPrimeFactor)
				}
			}

			if factor, more := <-primeFactorisationCh; more {
				t.Errorf("Received more prime factors than expected. Unexpected prime factor: %v", factor)
			}
		})
	}
}
