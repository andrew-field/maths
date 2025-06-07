package maths

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedResult int
	}{
		{[]int{0}, 0},
		{[]int{-10}, -10},
		{[]int{10}, 10},
		{[]int{-10, -10}, -10},
		{[]int{-10, -9}, -9},
		{[]int{-1, 0}, 0},
		{[]int{0, -1}, 0},
		{[]int{0, 0}, 0},
		{[]int{0, 1}, 1},
		{[]int{-1, 1}, 1},
		{[]int{1, 10}, 10},
		{[]int{-100, -100, -100}, -100},
		{[]int{-100, -100, -99}, -99},
		{[]int{-100, -1, 0}, 0},
		{[]int{1, 100, 10}, 100},
		{[]int{10, 5, 7, 2, 3, 6, 8, 1, 4, 9}, 10},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			// Check if the actual result matches the expected result.
			if actualResult := Max(tC.input...); actualResult != tC.expectedResult {
				t.Errorf("Expected result: %d, got result: %d", tC.expectedResult, actualResult)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	testCases := []struct {
		input, expectedResult int
		expectedError         error
	}{
		{math.MinInt + 1, math.MaxInt, nil},
		{-100, 100, nil},
		{-1, 1, nil},
		{0, 0, nil},
		{1, 1, nil},
		{100, 100, nil},
		{math.MaxInt, math.MaxInt, nil},
		{math.MinInt, 0, ErrAbsoluteValueOfMinInt},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %d", tC.input)
		t.Run(testName, func(t *testing.T) {
			actualResult, actualError := Abs(tC.input)

			// Check if the actual error matches the expected error.
			if !errors.Is(actualError, tC.expectedError) {
				t.Errorf("Expected error: %v, got error: %v", tC.expectedError, actualError)
			}

			// Check if the actual result matches the expected result.
			if actualResult != tC.expectedResult {
				t.Errorf("Expected result: %d, got result: %d", tC.expectedResult, actualResult)
			}
		})
	}
}

func ExampleAbs() {
	n := -10
	absValue, err := Abs(n)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Absolute value of", n, "is", absValue)
	}

	// Output: Absolute value of -10 is 10
}

func ExampleMax() {
	maxValue := Max(-100, -5, 0, 5, 10)
	fmt.Println("The maximum value is", maxValue)

	// Output: The maximum value is 10
}
