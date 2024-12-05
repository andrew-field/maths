package maths

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

var numberOfDivisorsTestCases = []struct {
	input, expectedResult int
}{
	{math.MinInt, 64},
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
	{math.MaxInt, 96},
}

var getDivisorsTestCases = []struct {
	input          int
	expectedResult []int
	expectedError  bool
}{
	{-4, []int{1, 2, 4}, false},
	{-3, []int{1, 3}, false},
	{-2, []int{1, 2}, false},
	{-1, []int{1}, false},
	{0, []int{}, false},
	{1, []int{1}, false},
	{2, []int{1, 2}, false},
	{3, []int{1, 3}, false},
	{4, []int{1, 2, 4}, false},
	{5, []int{1, 5}, false},
	{6, []int{1, 2, 3, 6}, false},
	{7, []int{1, 7}, false},
	{8, []int{1, 2, 4, 8}, false},
	{9, []int{1, 3, 9}, false},
	{10, []int{1, 2, 5, 10}, false},
	{100, []int{1, 2, 4, 5, 10, 20, 25, 50, 100}, false},
	{500, []int{1, 2, 4, 5, 10, 20, 25, 50, 100, 125, 250, 500}, false},
	{45664, []int{1, 2, 4, 8, 16, 32, 1427, 2854, 5708, 11416, 22832, 45664}, false},
	{7931265, []int{1, 3, 5, 15, 17, 19, 51, 57, 85, 95, 255, 285, 323, 969, 1615, 1637, 4845, 4911, 8185, 24555, 27829, 31103, 83487, 93309, 139145, 155515, 417435, 466545, 528751, 1586253, 2643755, 7931265}, false},
	{math.MaxInt, []int{1, 7, 49, 73, 127, 337, 511, 889, 2359, 3577, 6223, 9271, 16513, 24601, 42799, 64897, 92737, 172207, 299593, 454279, 649159, 649657, 1205449, 2097151, 3124327, 4544113, 4547599, 6769801,
		11777599, 21870289, 31252369, 31833193, 47388607, 47424961, 82443193, 82506439, 153092023, 218766583, 218934409, 331720249, 331974727, 577102351, 577545073, 859764727, 1531366081,
		1532540863, 2281422937, 2323823089, 3969050863, 4042815511, 6018353089, 6022970047, 10727786041, 15969960559, 15982211857, 27783356041, 27804669943, 42128471623, 42160790329, 60247241209,
		111789723913, 111875482999, 194483492287, 194632689601, 289740712999, 295125532303, 421730688463, 783128380993, 1362428827207, 2028184990993, 2029740905839, 2952114819241, 4398048608257,
		7651399633543, 14197294936951, 14208186340873, 20303320287433, 30786340257799, 53559797434801, 99457304386111, 142123242012031, 215504381804593, 374918582043607, 558552173248639, 994862694084217,
		1482142380982609, 2578521676503991, 3909865212740473, 10374996666878263, 18049651735527937, 27369056489183311, 72624976668147841, 126347562148695559, 188232082384791343, 1317624576693539401,
		9223372036854775807}, false},
	{math.MinInt, []int{}, true},
}

var sumOfDivisorsTestCases = []struct {
	input, expectedResult int
	expectedError         bool
}{
	{-3, 4, false},
	{-2, 3, false},
	{-1, 1, false},
	{0, 0, false},
	{1, 1, false},
	{2, 3, false},
	{3, 4, false},
	{4, 7, false},
	{5, 6, false},
	{6, 12, false},
	{7, 8, false},
	{8, 15, false},
	{9, 13, false},
	{10, 18, false},
	{100, 217, false},
	{500, 1092, false},
	{45664, 89964, false},
	{7931265, 14152320, false},
	{math.MinInt, 0, true}, // Can't get absolute value.
	{math.MaxInt, 0, true}, // Overflow error.
}

func TestNumberOfDivisors(t *testing.T) {
	testNumberOfDivisorsHelper(t, NumberOfDivisors)
}

func testNumberOfDivisorsHelper(t *testing.T, f func(int) int) {
	for _, tC := range numberOfDivisorsTestCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			// Check if the actual result matches the expected result.
			if actualResult := f(tC.input); actualResult != tC.expectedResult {
				t.Errorf("Expected number of divisors: %d, actual number of divisors: %d", tC.expectedResult, actualResult)
			}
		})
	}
}

func TestGetDivisors(t *testing.T) {
	testGetDivisorsHelper(t, GetDivisors)
}

func testGetDivisorsHelper(t *testing.T, f func(int) (<-chan int, error)) {
	for _, tC := range getDivisorsTestCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			divisorCh, actualError := f(tC.input)

			// Check if an error was returned and matches if an error was expected.
			if gotError := actualError != nil; gotError != tC.expectedError {
				t.Errorf("Expected error: %t, got error: %t, error: %v", tC.expectedError, gotError, actualError)
			}

			var actualDivisors []int
			for div := range divisorCh {
				actualDivisors = append(actualDivisors, div)
			}
			slices.Sort(actualDivisors)

			// Check if the each actual divisors match the expected divisors.
			if !slices.Equal(actualDivisors, tC.expectedResult) {
				t.Errorf("Actual divisors: %v. Expected divisors: %v.", actualDivisors, tC.expectedResult)
			}
		})
	}
}

func TestSumOfDivisors(t *testing.T) {
	testSumOfDivisorsHelper(t, SumOfDivisors)
}

func testSumOfDivisorsHelper(t *testing.T, f func(int) (int, error)) {
	for _, tC := range sumOfDivisorsTestCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			actualResult, actualError := f(tC.input)

			// Check if an error was returned and matches if an error was expected.
			if gotError := actualError != nil; gotError != tC.expectedError {
				t.Errorf("Expected error: %t, got error: %t, error: %v", tC.expectedError, gotError, actualError)
			}

			// Check if the actual result matches the expected result.
			if actualResult != tC.expectedResult {
				t.Errorf("Expected sum of divisors: %d, actual sum of divisors: %d", tC.expectedResult, actualResult)
			}
		})
	}
}
