package maths

import (
	"fmt"
	"math"
	"math/big"
	"slices"
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
		{new(big.Int).Exp(big.NewInt(-10), big.NewInt(35), nil), 36},
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
		{new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil), 21},
		{new(big.Int).Exp(big.NewInt(10), big.NewInt(35), nil), 36},
		{new(big.Int).Exp(big.NewInt(10), big.NewInt(98), nil), 99},
		{new(big.Int).Exp(big.NewInt(10), big.NewInt(99), nil), 100},
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
		{math.MinInt, []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 8}},
		{-10, []int{1, 0}},
		{-9, []int{9}},
		{-1, []int{1}},
		{0, []int{0}},
		{1, []int{1}},
		{9, []int{9}},
		{10, []int{1, 0}},
		{99, []int{9, 9}},
		{100, []int{1, 0, 0}},
		{500, []int{5, 0, 0}},
		{4563198, []int{4, 5, 6, 3, 1, 9, 8}},
		{math.MaxInt, []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7}},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			if digits := GetDigits(tC.input); !slices.Equal(digits, tC.expectedDigits) {
				t.Errorf("Actual digits: %v Expected digits: %v", digits, tC.expectedDigits)
			}
		})
	}
}

func TestGetDigitsBig(t *testing.T) {
	testCases := []struct {
		input          *big.Int
		expectedDigits []int
	}{
		{new(big.Int).Exp(big.NewInt(-10), big.NewInt(21), nil), []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{big.NewInt(math.MinInt), []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 8}},
		{big.NewInt(-10), []int{1, 0}},
		{big.NewInt(-9), []int{9}},
		{big.NewInt(-1), []int{1}},
		{big.NewInt(0), []int{0}},
		{big.NewInt(1), []int{1}},
		{big.NewInt(9), []int{9}},
		{big.NewInt(10), []int{1, 0}},
		{big.NewInt(99), []int{9, 9}},
		{big.NewInt(100), []int{1, 0, 0}},
		{big.NewInt(500), []int{5, 0, 0}},
		{big.NewInt(4563198), []int{4, 5, 6, 3, 1, 9, 8}},
		{big.NewInt(math.MaxInt), []int{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7}},
		{new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil), []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{new(big.Int).Exp(big.NewInt(2), big.NewInt(100), nil), []int{1, 2, 6, 7, 6, 5, 0, 6, 0, 0, 2, 2, 8, 2, 2, 9, 4, 0, 1, 4, 9, 6, 7, 0, 3, 2, 0, 5, 3, 7, 6}},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			if digits := GetDigitsBig(tC.input); !slices.Equal(digits, tC.expectedDigits) {
				t.Errorf("Actual digits: %v Expected digits: %v", digits, tC.expectedDigits)
			}
		})
	}
}
