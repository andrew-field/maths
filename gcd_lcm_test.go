package maths

import (
	"math/big"
	"testing"
)

func TestGCD(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedResult int
	}{
		{[]int{-6, 10, 0}, 2},
		{[]int{-12, -6, -4}, 2},
		{[]int{-2, 3, 4}, 1},
		{[]int{-3, -3}, 3},
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
		{[]int{3, 3}, 3},
		{[]int{2, 3, 4}, 1},
		{[]int{12, 6, 4}, 2},
		{[]int{6, 9}, 3},
		{[]int{6, 10}, 2},
		{[]int{6, 10, 0}, 2},
		{[]int{130, 65, 10}, 5},
		{[]int{4950, 3750, 450, 225}, 75},
		{[]int{527592, 91, 455}, 13},
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
		{[]int{-3, -3}, 3},
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
		{[]int{3, 3}, 3},
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

func TestLCMBig(t *testing.T) {
	testCases := []struct {
		input          []*big.Int
		expectedResult *big.Int
	}{
		{[]*big.Int{big.NewInt(-2), big.NewInt(-3), big.NewInt(-4), big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(-2), big.NewInt(-3), big.NewInt(-4)}, big.NewInt(12)},
		{[]*big.Int{big.NewInt(-3), big.NewInt(-3)}, big.NewInt(3)},
		{[]*big.Int{big.NewInt(-1), big.NewInt(2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(0), big.NewInt(-1)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(-2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(-1)}, big.NewInt(1)},
		{[]*big.Int{}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(1)}, big.NewInt(1)},
		{[]*big.Int{big.NewInt(2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(0), big.NewInt(1)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(1), big.NewInt(2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(3), big.NewInt(3)}, big.NewInt(3)},
		{[]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(4)}, big.NewInt(12)},
		{[]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(6), big.NewInt(10)}, big.NewInt(30)},
		{[]*big.Int{big.NewInt(5), big.NewInt(10), big.NewInt(65)}, big.NewInt(130)},
		{[]*big.Int{big.NewInt(75), big.NewInt(330), big.NewInt(225), big.NewInt(450)}, big.NewInt(4950)},
		{[]*big.Int{big.NewInt(13), big.NewInt(89), big.NewInt(456)}, big.NewInt(527592)},
	}

	for _, tC := range testCases {
		if actualResult := LCMBig(tC.input...); actualResult.Cmp(tC.expectedResult) != 0 {
			t.Errorf("Input in test: %v. Actual result: %v. Expected result: %v.", tC.input, actualResult, tC.expectedResult)
		}
	}
}
