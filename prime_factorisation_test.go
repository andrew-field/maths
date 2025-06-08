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

func FuzzPrimeFactorisation(f *testing.F) {
	testcases := []int{2, 5, 10, 20, 50, 100}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig int) {
		resultCh := PrimeFactorisation(orig)

		result := 1
		for factor := range resultCh {
			for range factor.Index {
				result *= factor.Value
			}
		}
		absOrig, err := Abs(orig)
		if err != nil {
			t.Skipf("Failed to get Abs of %d: %v", orig, err)
		}
		if result != absOrig {
			t.Errorf("Expected %d, got %d", absOrig, result)
		}
	})
}

func ExamplePrimeFactorisation() {
	input := 360
	resultCh := PrimeFactorisation(input)

	fmt.Printf("Prime factorisation of %d: ", input)
	for factor := range resultCh {
		fmt.Printf("%d^%d ", factor.Value, factor.Index)
	}

	// Output: Prime factorisation of 360: 2^3 3^2 5^1
}

func BenchmarkPrimeFactorisation(b *testing.B) {
	inputs := []int{10, 200, 3000, 40000, 500000, 6000000}
	for _, input := range inputs {
		b.Run(fmt.Sprintf("Input: %d", input), func(b *testing.B) {
			for b.Loop() {
				for range PrimeFactorisation(input) { // Just iterating through the channel to benchmark the function.
				}
			}
		})
	}
}
