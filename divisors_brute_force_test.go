package maths

import "testing"

func TestNumberOfDivisors2(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult int
	}{
		{-3, 2},
		{-2, 2},
		{-1, 1},
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 2},
		{6, 4},
		{7, 2},
		{8, 4},
		{9, 3},
		{10, 4},
		{100, 9},
		{500, 12},
		{45664, 12},
		{7931265, 32},
	}

	for _, tC := range testCases {
		if actualResult := NumberOfDivisors2(tC.input); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual number of divisors: %v. Expected number of divisors: %v.", tC.input, actualResult, tC.expectedResult)
		}
	}
}

func TestSumOfDivisors2(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult int
	}{
		{-3, 4},
		{-2, 3},
		{-1, 1},
		{0, 0},
		{1, 1},
		{2, 3},
		{3, 4},
		{4, 7},
		{5, 6},
		{6, 12},
		{7, 8},
		{8, 15},
		{9, 13},
		{10, 18},
		{100, 217},
		{500, 1092},
		{45664, 89964},
		{7931265, 14152320},
	}

	for _, tC := range testCases {
		if actualResult := SumOfDivisors2(tC.input); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual sum of divisors: %v. Expected sum of divisors: %v.", tC.input, actualResult, tC.expectedResult)
		}
	}
}

func TestDivisors2(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult []int
	}{
		{-4, []int{1, 4, 2}},
		{-3, []int{1, 3}},
		{-2, []int{1, 2}},
		{-1, []int{1}},
		{0, []int{}},
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 3}},
		{4, []int{1, 4, 2}},
		{5, []int{1, 5}},
		{6, []int{1, 6, 2, 3}},
		{7, []int{1, 7}},
		{8, []int{1, 8, 2, 4}},
		{9, []int{1, 9, 3}},
		{10, []int{1, 10, 2, 5}},
		{100, []int{1, 100, 2, 50, 4, 25, 5, 20, 10}},
		{500, []int{1, 500, 2, 250, 4, 125, 5, 100, 10, 50, 20, 25}},
		{45664, []int{1, 45664, 2, 22832, 4, 11416, 8, 5708, 16, 2854, 32, 1427}},
		{7931265, []int{1, 7931265, 3, 2643755, 5, 1586253, 15, 528751, 17, 466545, 19, 417435, 51, 155515, 57, 139145, 85, 93309, 95, 83487, 255, 31103, 285, 27829, 323, 24555, 969, 8185, 1615, 4911, 1637, 4845}},
	}

	for _, tC := range testCases {
		divisorChannel := Divisors2(tC.input)

		for _, expectedResult := range tC.expectedResult {
			if actualResult := <-divisorChannel; actualResult != expectedResult {
				t.Errorf("Input in test: %v. Actual divisor: %v. Expected divisor: %v.", tC.input, actualResult, expectedResult)
			}
		}

		// Check the divisor channel does not have too many values.
		if divisor, more := <-divisorChannel; more {
			t.Errorf("Received more divisors than expected. Input in test: %v. Unexpected divisor: %v", tC.input, divisor)
		}
	}
}
