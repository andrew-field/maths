package maths

import "testing"

func TestMaxPath(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedResult int
	}{
		{[]int{1, 2, -3}, 3},
		{[]int{-1}, -1},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3}, 4},
		{[]int{1, 2, 3, 4, 5}, 8},
		{[]int{75, 95, 64, 17, 47, 82, 18, 35, 87, 10}, 274},
		{[]int{75, 95, 64, 17, 47, 82, 18, 35, 87, 10, 24, 12, 54, 38, 20, 46, 35, 42, 64, 21, 45}, 338},
	}

	// Uses CreateTree to create the trees with which to test the function.
	for _, tC := range testCases {
		tree := CreateTree(tC.input...)
		if actualResult := MaxPath(tree); actualResult != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual result: %v. Expected result: %v", tC.input, actualResult, tC.expectedResult)
		}
	}
}

// func TestGetMaximumPathSumOfPyramidUsingMaximumSlotsANDGetMaximumPathSumOfPyramidUsingRecursiveFunction(t *testing.T) {
// 	testCases := []struct {
// 		input          [][]float64
// 		expectedResult float64
// 	}{
// 		{[][]float64{{0}}, 0},
// 		{[][]float64{{1}}, 1},
// 		{[][]float64{{1}, {2, 3}}, 4},
// 		{[][]float64{{1}, {2, 3}, {4, 5, 6}}, 10},
// 		{[][]float64{{75}, {95, 64}, {17, 47, 82}, {18, 35, 87, 10}}, 308},
// 		{[][]float64{{75}, {95, 64}, {17, 47, 82}, {18, 35, 87, 10}, {24, 12, 54, 38, 20}, {46, 35, 42, 64, 21, 45}}, 426},
// 	}
// 	type functionToTest struct {
// 		name     string
// 		function func([][]float64) float64
// 	}
// 	for _, tC := range testCases {
// 		functions := []functionToTest{{"GetMaximumPathSumOfPyramidUsingMaximumSlots", GetMaximumPathSumOfPyramidUsingMaximumSlots},
// 			{"GetMaximumPathSumOfPyramidUsingRecursiveFunction", GetMaximumPathSumOfPyramidUsingRecursiveFunction}}
// 		for _, function := range functions {
// 			if actualResult := function.function(tC.input); tC.expectedResult != actualResult {
// 				t.Errorf("%v has failed. Input in test: %v. Expected result: %v. Actual result: %v.", function.name, tC.input, tC.expectedResult, actualResult)
// 			}
// 		}
// 	}

// 	// Testing that the code panics if there is a bad input.
// 	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingMaximumSlots, "GetMaximumPathSumOfPyramidUsingMaximumSlots", nil)
// 	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingMaximumSlots, "GetMaximumPathSumOfPyramidUsingMaximumSlots", [][]float64{{1, 2}})
// 	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingMaximumSlots, "GetMaximumPathSumOfPyramidUsingMaximumSlots", [][]float64{{1}, {2, 3, 4}})
// 	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingMaximumSlots, "GetMaximumPathSumOfPyramidUsingMaximumSlots", make([][]float64, 0))
// 	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingRecursiveFunction, "GetMaximumPathSumOfPyramidUsingRecursiveFunction", nil)
// 	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingRecursiveFunction, "GetMaximumPathSumOfPyramidUsingRecursiveFunction", [][]float64{{1, 2}})
// 	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingRecursiveFunction, "GetMaximumPathSumOfPyramidUsingRecursiveFunction", [][]float64{{1}, {2, 3, 4}})
// 	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingRecursiveFunction, "GetMaximumPathSumOfPyramidUsingRecursiveFunction", make([][]float64, 0))

// }

// func CheckGetMaximumPathSumOfPyramidPanics(t *testing.T, function func([][]float64) float64, functionName string, inputForFunction [][]float64) {
// 	defer func() {
// 		if r := recover(); r == nil {
// 			t.Errorf("%v has failed. The code did not panic.", functionName)
// 		}
// 	}()
// 	function(inputForFunction)
// }
