package maths

import (
	"math"
	"math/big"
	"testing"
)

func TestNumberOfDigits(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult int
	}{
		{math.MinInt64, 19},
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
		{math.MaxInt64, 19},
	}

	for _, tC := range testCases {
		if actualNumberOfDigits := NumberOfDigits(tC.input); actualNumberOfDigits != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual number of digits: %v. Expected number of digits: %v.", tC.input, actualNumberOfDigits, tC.expectedResult)
		}
	}
}

func TestNumberOfDigitsBig(t *testing.T) {
	testCases := []struct {
		input          *big.Int
		expectedResult int
	}{
		{big.NewInt(0).Exp(big.NewInt(-10), big.NewInt(35), nil), 36},
		{big.NewInt(math.MinInt64), 19},
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
		{big.NewInt(math.MaxInt64), 19},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(20), nil), 21},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(35), nil), 36},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(98), nil), 99},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(99), nil), 100},
	}

	for _, tC := range testCases {
		if actualNumberOfDigits := NumberOfDigitsBig(tC.input); actualNumberOfDigits != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual number of digits: %v. Expected number of digits: %v.", tC.input, actualNumberOfDigits, tC.expectedResult)
		}
	}
}

func TestDigits(t *testing.T) {
	testCases := []struct {
		input          int
		expectedDigits []int
	}{
		{math.MinInt64, []int{8, 0, 8, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9}},
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
		{math.MaxInt64, []int{7, 0, 8, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9}},
	}

	for _, tC := range testCases {
		digitChannel := Digits(tC.input)

		checkDigitsAreCorrect(tC, digitChannel, t)
	}
}

func TestDigitsBig(t *testing.T) {
	testCases := []struct {
		input          *big.Int
		expectedDigits []int
	}{
		{big.NewInt(0).Exp(big.NewInt(-10), big.NewInt(21), nil), []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{big.NewInt(math.MinInt64), []int{8, 0, 8, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9}},
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
		{big.NewInt(math.MaxInt64), []int{7, 0, 8, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9}},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(20), nil), []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{big.NewInt(0).Exp(big.NewInt(2), big.NewInt(100), nil), []int{6, 7, 3, 5, 0, 2, 3, 0, 7, 6, 9, 4, 1, 0, 4, 9, 2, 2, 8, 2, 2, 0, 0, 6, 0, 5, 6, 7, 6, 2, 1}},
	}

	for _, tC := range testCases {
		digitChannel := DigitsBig(tC.input)

		checkDigitsAreCorrect(tC, digitChannel, t)
	}
}

func checkDigitsAreCorrect[T int | *big.Int](tC struct {
	input          T
	expectedDigits []int
}, digitChannel <-chan int, t *testing.T) {
	for _, expectedDigit := range tC.expectedDigits {
		if actualDigit := <-digitChannel; actualDigit != expectedDigit {
			t.Errorf("Input in test: %v. Actual digit: %v. Expected digit: %v.", tC.input, actualDigit, expectedDigit)
		}
	}

	// Check the digit channel does not have too many values.
	if digit, more := <-digitChannel; more {
		t.Errorf("Received more digits than expected. Input in test: %v. Unexpected digit: %v", tC.input, digit)
	}
}
