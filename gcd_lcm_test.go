package maths

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestGCD(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedResult int
		expectedError  bool
	}{
		{[]int{-6, 10, 0}, 2, false},
		{[]int{-12, -6, -4}, 2, false},
		{[]int{-2, 3, 4}, 1, false},
		{[]int{-3, -3}, 3, false},
		{[]int{-1, 2}, 1, false},
		{[]int{0, -1}, 1, false},
		{[]int{-2}, 2, false},
		{[]int{-1}, 1, false},
		{[]int{1}, 1, false},
		{[]int{2}, 2, false},
		{[]int{0, 1}, 1, false},
		{[]int{1, 2}, 1, false},
		{[]int{3, 3}, 3, false},
		{[]int{2, 3, 4}, 1, false},
		{[]int{12, 6, 4}, 2, false},
		{[]int{6, 9}, 3, false},
		{[]int{6, 10}, 2, false},
		{[]int{6, 10, 0}, 2, false},
		{[]int{130, 65, 10}, 5, false},
		{[]int{4950, 3750, 450, 225}, 75, false},
		{[]int{527592, 91, 455}, 13, false},
		{[]int{math.MaxInt}, math.MaxInt, false},
		{[]int{}, 0, false},
		{[]int{0}, 0, false},
		{[]int{0, 0, 0}, 0, false},
		{[]int{0, 0, 0, 10}, 10, false},
		{[]int{math.MinInt}, 0, true},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			actualResult, actualError := GCD(tC.input...)

			// Check if an error was returned and matches if an error was expected.
			if gotError := actualError != nil; gotError != tC.expectedError {
				t.Errorf("Expected error: %t, got error: %t, error: %v", tC.expectedError, gotError, actualError)
			}

			// Check if the actual result matches the expected result.
			if actualResult != tC.expectedResult {
				t.Errorf("Expected GCD: %d, actual GCD: %d", tC.expectedResult, actualResult)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedResult int
		expectedError  bool
	}{
		{[]int{-2, -3, -4}, 12, false},
		{[]int{-3, -3}, 3, false},
		{[]int{-1, 2}, 2, false},
		{[]int{-1, -1}, 1, false},
		{[]int{-2}, 2, false},
		{[]int{-1}, 1, false},
		{[]int{}, 0, false},
		{[]int{0}, 0, false},
		{[]int{0, 0, 0}, 0, false},
		{[]int{1}, 1, false},
		{[]int{2}, 2, false},
		{[]int{1, 2}, 2, false},
		{[]int{3, 3}, 3, false},
		{[]int{2, 3, 4}, 12, false},
		{[]int{6, 10}, 30, false},
		{[]int{5, 10, 65}, 130, false},
		{[]int{75, 330, 225, 450}, 4950, false},
		{[]int{13, 89, 456}, 527592, false},
		{[]int{-2, -3, 4, 0}, 0, false},
		{[]int{2, 3, 4, 0}, 0, false},
		{[]int{1, 0}, 0, false},
		{[]int{math.MinInt, 2}, 0, true},         // Fail GCD.
		{[]int{math.MinInt}, 0, true},            // Fail Abs.
		{[]int{math.MaxInt}, math.MaxInt, false}, // Overflow.
		{[]int{math.MaxInt, 2}, 0, true},         // Overflow.
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			actualResult, actualError := LCM(tC.input...)

			// Check if an error was returned and matches if an error was expected.
			if gotError := actualError != nil; gotError != tC.expectedError {
				t.Errorf("Expected error: %t, got error: %t, error: %v", tC.expectedError, gotError, actualError)
			}

			// Check if the actual result matches the expected result.
			if actualResult != tC.expectedResult {
				t.Errorf("Expected LCM: %d, actual LCM: %d", tC.expectedResult, actualResult)
			}
		})
	}
}

func TestLCMBig(t *testing.T) {
	testCases := []struct {
		input          []*big.Int
		expectedResult *big.Int
	}{
		{[]*big.Int{big.NewInt(-2), big.NewInt(-3), big.NewInt(-4)}, big.NewInt(12)},
		{[]*big.Int{big.NewInt(-3), big.NewInt(-3)}, big.NewInt(3)},
		{[]*big.Int{big.NewInt(-1), big.NewInt(2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(-1), big.NewInt(-1)}, big.NewInt(1)},
		{[]*big.Int{big.NewInt(-2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(-1)}, big.NewInt(1)},
		{[]*big.Int{}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(1)}, big.NewInt(1)},
		{[]*big.Int{big.NewInt(2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(1), big.NewInt(2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(3), big.NewInt(3)}, big.NewInt(3)},
		{[]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(4)}, big.NewInt(12)},
		{[]*big.Int{big.NewInt(6), big.NewInt(10)}, big.NewInt(30)},
		{[]*big.Int{big.NewInt(5), big.NewInt(10), big.NewInt(65)}, big.NewInt(130)},
		{[]*big.Int{big.NewInt(75), big.NewInt(330), big.NewInt(225), big.NewInt(450)}, big.NewInt(4950)},
		{[]*big.Int{big.NewInt(13), big.NewInt(89), big.NewInt(456)}, big.NewInt(527592)},
		{[]*big.Int{big.NewInt(-2), big.NewInt(-3), big.NewInt(4), big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(1), big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(math.MinInt), big.NewInt(2)}, big.NewInt(0).Abs(big.NewInt(math.MinInt))},
		{[]*big.Int{big.NewInt(math.MinInt), big.NewInt(0).Abs(big.NewInt(math.MinInt))}, big.NewInt(0).Abs(big.NewInt(math.MinInt))},
		{[]*big.Int{big.NewInt(math.MaxInt), big.NewInt(2)}, big.NewInt(0).Mul(big.NewInt(math.MaxInt), big.NewInt(2))},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			// Check if the actual result matches the expected result.
			if actualResult := LCMBig(tC.input...); actualResult.Cmp(tC.expectedResult) != 0 {
				t.Errorf("Expected LCM: %v, actual LCM: %v", tC.expectedResult, actualResult) // Can print the big.Int values OK.
			}
		})
	}
}
