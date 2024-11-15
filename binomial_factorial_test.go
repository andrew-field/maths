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

			// Check if an error was returned and matches if an error was expected.
			if gotError := actualError != nil; gotError != tC.expectedError {
				t.Errorf("Expected error: %t, got error: %t, error: %v", tC.expectedError, gotError, actualError)
			}

			// Check if the actual result matches the expected result
			if actualResult != tC.expectedResult {
				t.Errorf("Expected result: %d, got result: %d", tC.expectedResult, actualResult)
			}
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

			// Check if an error was returned and matches if an error was expected.
			if gotError := actualError != nil; gotError != tC.expectedError {
				t.Errorf("Expected error: %t, got error: %t, error: %v", tC.expectedError, gotError, actualError)
			}

			// Check if the actual result matches the expected result
			if actualResult != tC.expectedResult {
				t.Errorf("Expected result: %d, got result: %d", tC.expectedResult, actualResult)
			}
		})
	}
}
