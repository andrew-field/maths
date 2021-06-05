package maths

import (
	"testing"
)

func TestNumberOfDivisors(t *testing.T) {
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
		if actualResult := NumberOfDivisors(tC.input); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual number of divisors: %v. Expected number of divisors: %v.", tC.input, actualResult, tC.expectedResult)
		}
	}
}

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

func TestSumOfDivisors(t *testing.T) {
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
		if actualResult := SumOfDivisors(tC.input); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual sum of divisors: %v. Expected sum of divisors: %v.", tC.input, actualResult, tC.expectedResult)
		}
	}
}

func TestDivisors(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult []int
	}{
		{-4, []int{1, 2, 4}},
		{-3, []int{1, 3}},
		{-2, []int{1, 2}},
		{-1, []int{1}},
		{0, []int{}},
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 3}},
		{4, []int{1, 2, 4}},
		{5, []int{1, 5}},
		{6, []int{1, 2, 3, 6}},
		{7, []int{1, 7}},
		{8, []int{1, 2, 4, 8}},
		{9, []int{1, 3, 9}},
		{10, []int{1, 2, 5, 10}},
		{100, []int{1, 2, 4, 5, 10, 20, 25, 50, 100}},
		{500, []int{1, 2, 4, 5, 10, 20, 25, 50, 100, 125, 250, 500}},
		{45664, []int{1, 2, 4, 8, 16, 32, 1427, 2854, 5708, 11416, 22832, 45664}},
		{7931265, []int{1, 3, 5, 15, 17, 51, 85, 255, 19, 57, 95, 285, 323, 969, 1615, 4845, 1637, 4911, 8185, 24555, 27829, 83487, 139145, 417435, 31103, 93309, 155515, 466545, 528751, 1586253, 2643755, 7931265}},
	}

	for _, tC := range testCases {
		divisorChannel := Divisors(tC.input)

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
