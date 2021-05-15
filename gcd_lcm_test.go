package maths

import "testing"

func TestGCD(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedResult int
	}{
		{[]int{-6, 10, 0}, 2},
		{[]int{-12, -3, -4}, 2},
		{[]int{-2, 3, 4}, 1},
		{[]int{-1, 2}, 1},
		{[]int{0, -1}, 1},
		{[]int{-2}, 2},
		{[]int{-1}, 1},
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{2}, 2},
		{[]int{0, 1}, 1},
		{[]int{1, 2}, 1},
		{[]int{2, 3, 4}, 1},
		{[]int{12, 3, 4}, 2},
		{[]int{6, 9}, 3},
		{[]int{6, 10}, 2},
		{[]int{6, 10, 0}, 2},
		{[]int{130, 65, 10}, 5},
		{[]int{4950, 330, 450, 225}, 75},
		{[]int{527592, 89, 456}, 13},
	}

	for _, tC := range testCases {
		if actualResult := GCD(tC.input...); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual result: %v. Expected result: %v", tC.input, actualResult, tC.expectedResult)
		}
	}
}

func TestLCM(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedResult int
	}{
		{[]int{-2, -3, -4, 0}, 0},
		{[]int{-2, -3, -4}, 12},
		{[]int{-1, 2}, 2},
		{[]int{0, -1}, 0},
		{[]int{-2}, 2},
		{[]int{-1}, 1},
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{2}, 2},
		{[]int{0, 1}, 0},
		{[]int{1, 2}, 2},
		{[]int{2, 3, 4}, 12},
		{[]int{2, 3, 4, 0}, 0},
		{[]int{6, 10}, 30},
		{[]int{5, 10, 65}, 130},
		{[]int{75, 330, 225, 450}, 4950},
		{[]int{13, 89, 456}, 527592},
	}

	for _, tC := range testCases {
		if actualResult := LCM(tC.input...); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual result: %v. Expected result: %v.", tC.input, actualResult, tC.expectedResult)
		}
	}
}
