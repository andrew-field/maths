package maths

import (
	"fmt"
	"slices"
	"testing"
)

func TestNumberOfDivisorsBruteForce(t *testing.T) {
	for _, tC := range numberOfDivisorsTestCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			// Check if the actual result matches the expected result.
			if actualResult := NumberOfDivisorsBruteForce(tC.input); actualResult != tC.expectedResult {
				t.Errorf("Expected number of divisors: %d, actual number of divisors: %d", tC.expectedResult, actualResult)
			}
		})
	}
}

func TestGetDivisorsBruteForce(t *testing.T) {
	for _, tC := range getDivisorsTestCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			divisorCh, actualError := GetDivisorsBruteForce(tC.input)

			// Check if an error was returned and matches if an error was expected.
			if gotError := actualError != nil; gotError != tC.expectedError {
				t.Errorf("Expected error: %t, got error: %t, error: %v", tC.expectedError, gotError, actualError)
			}

			var actualDivisors []int
			for div := range divisorCh {
				actualDivisors = append(actualDivisors, div)
			}
			slices.Sort(actualDivisors)

			// Check if the each actual divisors match the expected divisors.
			if !slices.Equal(actualDivisors, tC.expectedResult) {
				t.Errorf("Actual divisors: %v. Expected divisors: %v.", actualDivisors, tC.expectedResult)
			}
		})
	}
}

func TestSumOfDivisorsBruteForce(t *testing.T) {
	for _, tC := range sumOfDivisorsTestCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			actualResult, actualError := SumOfDivisorsBruteForce(tC.input)

			// Check if an error was returned and matches if an error was expected.
			if gotError := actualError != nil; gotError != tC.expectedError {
				t.Errorf("Expected error: %t, got error: %t, error: %v", tC.expectedError, gotError, actualError)
			}

			// Check if the actual result matches the expected result.
			if actualResult != tC.expectedResult {
				t.Errorf("Expected sum of divisors: %d, actual sum of divisors: %d", tC.expectedResult, actualResult)
			}
		})
	}
}
