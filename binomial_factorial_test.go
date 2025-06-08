package maths

import (
	"fmt"
	"math"
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input, expectedResult int
		expectedError         bool
	}{
		{-10, 3628800, false},
		{-2, 2, false},
		{-1, 1, false},
		{0, 1, false},
		{1, 1, false},
		{2, 2, false},
		{3, 6, false},
		{10, 3628800, false},
		{1000, 0, true},
		{math.MinInt, 0, true},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			actualResult, actualError := Factorial(tC.input)

			checkResults(t, tC.expectedResult, tC.expectedError, actualResult, actualError)
		})
	}
}

func TestBinomial(t *testing.T) {
	testCases := []struct {
		n, k, expectedResult int
		expectedError        bool
	}{
		{10, -5, 252, false},
		{-2, -1, 2, false},
		{-1, 0, 1, false},
		{0, 0, 1, false},
		{1, 0, 1, false},
		{1, 1, 1, false},
		{2, 1, 2, false},
		{2, 2, 1, false},
		{10, 10, 1, false},
		{10, 5, 252, false},
		{math.MinInt, 1, 0, true},
		{1, math.MinInt, 0, true},
		{5, 6, 0, true},
		{-5, -6, 0, true},
		{1000, 1, 0, true},
		// Can not currently test the case where fact(k) returns an error. If this were to return an error, then fact(absN) would have already thrown an error earlier.
		// Can not currently test the case where fact(differenceOfAbsolute) returns an error. If this were to return an error, then fact(absN) would have already thrown an error earlier.
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: n:%d, k:%d", tC.n, tC.k)
		t.Run(testName, func(t *testing.T) {
			actualResult, actualError := Binomial(tC.n, tC.k)

			checkResults(t, tC.expectedResult, tC.expectedError, actualResult, actualError)
		})
	}
}

func checkResults(t testing.TB, expectedResult int, expectedError bool, actualResult int, actualError error) {
	t.Helper()
	// Check if an error was returned and matches if an error was expected.
	if gotError := actualError != nil; gotError != expectedError {
		t.Errorf("Expected error: %t, got error: %t, error: %v", expectedError, gotError, actualError)
	}

	// Check if the actual result matches the expected result.
	if actualResult != expectedResult {
		t.Errorf("Expected result: %d, got result: %d", expectedResult, actualResult)
	}
}

func ExampleFactorial() {
	n := 10
	p, err := Factorial(n)
	if err != nil {
		fmt.Printf("Error calculating the factorial of %d: %v\n", n, err)
	} else {
		fmt.Printf("The factorial of %d is %d\n", n, p)
	}

	n = 21
	p, err = Factorial(n)
	if err != nil {
		fmt.Printf("Error calculating the factorial of %d: %v\n", n, err)
	} else {
		fmt.Printf("The factorial of %d is %d\n", n, p)
	}

	// Output:
	// The factorial of 10 is 3628800
	// Error calculating the factorial of 21: failed to get fact(21): failed to calculate 21 * 2432902008176640000. The result is too large to hold in an int variable: arithmetic overflow detected
}

func ExampleBinomial() {
	n, k := 10, 3
	p, err := Binomial(n, k)
	if err != nil {
		fmt.Printf("Error calculating the binomial coefficient of %d choose %d: %v\n", n, k, err)
	} else {
		fmt.Printf("The binomial coefficient of %d choose %d is %d\n", n, k, p)
	}

	n, k = 22, 5
	p, err = Binomial(n, k)
	if err != nil {
		fmt.Printf("Error calculating the binomial coefficient of %d choose %d: %v\n", n, k, err)
	} else {
		fmt.Printf("The binomial coefficient of %d choose %d is %d\n", n, k, p)
	}

	// Output:
	// The binomial coefficient of 10 choose 3 is 120
	// Error calculating the binomial coefficient of 22 choose 5: failed to get fact(22): failed to calculate 21 * 2432902008176640000. The result is too large to hold in an int variable: arithmetic overflow detected
}
