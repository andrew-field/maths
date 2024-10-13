package maths

import "testing"

func TestNumberOfDivisorsBruteForce(t *testing.T) {
	for _, tC := range numberOfDivisorsTestCases {
		if actualResult := NumberOfDivisorsBruteForce(tC.input); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual number of divisors: %v. Expected number of divisors: %v.", tC.input, actualResult, tC.expectedResult)
		}
	}
}

func TestSumOfDivisorsBruteForce(t *testing.T) {
	for _, tC := range sumOfDivisorsTestCases {
		if actualResult := SumOfDivisorsBruteForce(tC.input); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual sum of divisors: %v. Expected sum of divisors: %v.", tC.input, actualResult, tC.expectedResult)
		}
	}
}

func TestGetDivisorsBruteForce(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult []int
	}{
		{-4, []int{1, 4, 2}},
		{-3, []int{1, 3}},
		{-2, []int{1, 2}},
		{-1, []int{1}},
		{1, []int{1}}, // This is above the line below so JSCPD linter doesn't see a clone with the test cases in TestGCD.
		{0, []int{}},
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
		divisorChannel := GetDivisorsBruteForce(tC.input)

		checkDivisorsAreCorrect(tC, divisorChannel, t)
	}
}
