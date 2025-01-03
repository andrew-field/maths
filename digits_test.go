package maths

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestNumberOfDigits(t *testing.T) {
	testCases := []struct {
		input, expectedResult int
	}{
		{math.MinInt, 19},
		{-10, 2},
		{-9, 1},
		{-1, 1},
		{0, 1},
		{1, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{500, 3},
		{4563198, 7},
		{math.MaxInt, 19},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			// Check if the actual result matches the expected result.
			if actualNumberOfDigits := NumberOfDigits(tC.input); actualNumberOfDigits != tC.expectedResult {
				t.Errorf("Actual number of digits: %d. Expected number of digits: %d", actualNumberOfDigits, tC.expectedResult)
			}
		})
	}
}

func TestNumberOfDigitsBig(t *testing.T) {
	testCases := []struct {
		input          *big.Int
		expectedResult int
	}{
		{big.NewInt(0).Exp(big.NewInt(-10), big.NewInt(35), nil), 36},
		{big.NewInt(math.MinInt), 19},
		{big.NewInt(-10), 2},
		{big.NewInt(-9), 1},
		{big.NewInt(-1), 1},
		{big.NewInt(0), 1},
		{big.NewInt(1), 1},
		{big.NewInt(9), 1},
		{big.NewInt(10), 2},
		{big.NewInt(99), 2},
		{big.NewInt(100), 3},
		{big.NewInt(500), 3},
		{big.NewInt(4563198), 7},
		{big.NewInt(math.MaxInt), 19},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(20), nil), 21},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(35), nil), 36},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(98), nil), 99},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(99), nil), 100},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			// Check if the actual result matches the expected result.
			if actualNumberOfDigits := NumberOfDigitsBig(tC.input); actualNumberOfDigits != tC.expectedResult {
				t.Errorf("Actual number of digits: %d. Expected number of digits: %d", actualNumberOfDigits, tC.expectedResult)
			}
		})
	}
}

func TestGetDigits(t *testing.T) {
	testCases := []struct {
		input          int
		expectedDigits []int
	}{
		{math.MinInt, []int{8, 0, 8, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9}},
		{-10, []int{0, 1}},
		{-9, []int{9}},
		{-1, []int{1}},
		{0, []int{0}},
		{1, []int{1}},
		{9, []int{9}},
		{10, []int{0, 1}},
		{99, []int{9, 9}},
		{100, []int{0, 0, 1}},
		{500, []int{0, 0, 5}},
		{4563198, []int{8, 9, 1, 3, 6, 5, 4}},
		{math.MaxInt, []int{7, 0, 8, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9}},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			digitCh := GetDigits(tC.input)

			checkDigitResults(t, tC.expectedDigits, digitCh)
		})
	}
}

func TestGetDigitsBig(t *testing.T) {
	testCases := []struct {
		input          *big.Int
		expectedDigits []int
	}{
		{big.NewInt(0).Exp(big.NewInt(-10), big.NewInt(21), nil), []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{big.NewInt(math.MinInt), []int{8, 0, 8, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9}},
		{big.NewInt(-10), []int{0, 1}},
		{big.NewInt(-9), []int{9}},
		{big.NewInt(-1), []int{1}},
		{big.NewInt(0), []int{0}},
		{big.NewInt(1), []int{1}},
		{big.NewInt(9), []int{9}},
		{big.NewInt(10), []int{0, 1}},
		{big.NewInt(99), []int{9, 9}},
		{big.NewInt(100), []int{0, 0, 1}},
		{big.NewInt(500), []int{0, 0, 5}},
		{big.NewInt(4563198), []int{8, 9, 1, 3, 6, 5, 4}},
		{big.NewInt(math.MaxInt), []int{7, 0, 8, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9}},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(20), nil), []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{big.NewInt(0).Exp(big.NewInt(2), big.NewInt(100), nil), []int{6, 7, 3, 5, 0, 2, 3, 0, 7, 6, 9, 4, 1, 0, 4, 9, 2, 2, 8, 2, 2, 0, 0, 6, 0, 5, 6, 7, 6, 2, 1}},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			digitCh := GetDigitsBig(tC.input)

			checkDigitResults(t, tC.expectedDigits, digitCh)
		})
	}
}

func checkDigitResults(t *testing.T, expectedDigits []int, digitCh <-chan int) {
	// Check if the each actual digit matches the expected digit.
	for index, expectedDigit := range expectedDigits {
		if actualDigit := <-digitCh; actualDigit != expectedDigit {
			t.Errorf("Actual digit: %d. Expected digit: %d at index %d", actualDigit, expectedDigit, index)
		}
	}

	// Check the digit channel does not have too many values.
	if digit, more := <-digitCh; more {
		t.Errorf("Received more digits than expected. Unexpected digit: %d", digit)
	}
}
