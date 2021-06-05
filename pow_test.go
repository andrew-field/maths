package maths

import (
	"testing"
)

func TestPow(t *testing.T) {
	testCases := []struct {
		inputX, inputY int
		expectedResult int
	}{
		{-10, 6, 1000000},
		{-10, -6, 1000000},
		{-10, 5, -100000},
		{-10, -5, -100000},
		{-1, 1, -1},
		{-1, 0, 1},
		{0, 1, 0},
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{2, 2, 4},
		{2, 3, 8},
		{3, 3, 27},
	}

	for _, tC := range testCases {
		if actualValue := Pow(tC.inputX, tC.inputY); actualValue != tC.expectedResult {
			t.Errorf("Input in test: %v, %v. Actual value: %v. Expected value: %v.", tC.inputX, tC.inputY, actualValue, tC.expectedResult)
		}
	}
}
