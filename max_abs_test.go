package maths

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedResult int
	}{
		{[]int{}, 0},
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
		if actualMax := Max(tC.input...); actualMax != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual max: %v. Expected max: %v.", tC.input, actualMax, tC.expectedResult)
		}
	}
}

func TestAbs(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult int
	}{
		{math.MinInt64 + 1, math.MaxInt64},
		{-100, 100},
		{-1, 1},
		{0, 0},
		{1, 1},
		{100, 100},
		{math.MaxInt64, math.MaxInt64},
	}

	for _, tC := range testCases {
		if actualValue := Abs(tC.input); actualValue != tC.expectedResult {
			t.Errorf("Input in test: %v. Actual value: %v. Expected value: %v.", tC.input, actualValue, tC.expectedResult)
		}
	}
}
