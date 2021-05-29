package maths

import (
	"testing"
)

func TestFact(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult int
	}{
		{-10, 3628800},
		{-2, 2},
		{-1, 1},
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{10, 3628800},
	}

	for _, tC := range testCases {
		if actualResult := Fact(tC.input); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual factorial: %v. Expected factorial: %v.", tC.input, actualResult, tC.expectedResult)
		}
	}
}

func TestBinomial(t *testing.T) {
	testCases := []struct {
		n, k           int
		expectedResult int
	}{
		{10, -5, 252},
		{-2, -1, 2},
		{-1, 0, 1},
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{2, 1, 2},
		{2, 2, 1},
		{2, 3, 0},
		{2, 4, 0},
		{10, 5, 252},
		{10, 10, 1},
		{10, 11, 0},
	}

	for _, tC := range testCases {
		if actualResult := Binomial(tC.n, tC.k); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v, %v. Actual result: %v. Expected result: %v.", tC.n, tC.k, actualResult, tC.expectedResult)
		}
	}
}
