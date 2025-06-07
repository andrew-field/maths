package maths

import (
	"fmt"
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	testCases := []struct {
		x, y, expectedResult int
		expectedError        bool
	}{
		{-10, 6, 1000000, false},
		{-10, -6, 1000000, false},
		{-10, 5, -100000, false},
		{-10, -5, -100000, false},
		{-1, 1, -1, false},
		{-1, 0, 1, false},
		{0, 1, 0, false},
		{0, 0, 1, false},
		{1, 0, 1, false},
		{1, 1, 1, false},
		{1, math.MinInt, 1, false},
		{2, 2, 4, false},
		{2, 3, 8, false},
		{3, 3, 27, false},
		{math.MaxInt, 1, math.MaxInt, false},
		{math.MinInt, 1, math.MinInt, false},
		{-2, 63, math.MinInt, false}, // Same an math.MinInt.
		{2, math.MinInt, 0, true},    // Fail to get Abs().
		{math.MaxInt, 2, 0, true},    // Overflow
		{2, 70, 0, true},             // Overflow
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: x:%d, y:%d", tC.x, tC.y)
		t.Run(testName, func(t *testing.T) {
			actualResult, actualError := Pow(tC.x, tC.y)

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

func ExamplePow() {
	m, n := 3, 5
	result, err := Pow(m, n)
	if err != nil {
		fmt.Printf("Error calculating %d^%d: %v", m, n, err)
	} else {
		fmt.Printf("%d^%d = %d", m, n, result)
	}

	// Output: 3^5 = 243
}
