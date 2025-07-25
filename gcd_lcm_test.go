package maths

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestGCD(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{[]int{-6, 10, 0}, 2},
		{[]int{-12, -6, -4}, 2},
		{[]int{-2, 3, 4}, 1},
		{[]int{-3, -3}, 3},
		{[]int{-1, 2}, 1},
		{[]int{0, -1}, 1},
		{[]int{-2}, 2},
		{[]int{-1}, 1},
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
		{[]int{math.MaxInt}, math.MaxInt},
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{0, 0, 0}, 0},
		{[]int{0, 0, 0, 10}, 10},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			got, gotError := GCD(tC.input...)

			checkResult(t, tC.want, got, gotError)
		})
	}

	errorTestCases := []struct {
		desc      string
		input     []int
		wantError error
	}{
		{"|math.MinInt| is too large to hold in an int", []int{math.MinInt}, ErrAbsoluteValueOfMinInt},
	}

	for _, tC := range errorTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, gotError := GCD(tC.input...)

			checkError(t, gotError, tC.wantError)
		})
	}
}

func TestLCM(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{[]int{-2, -3, -4}, 12},
		{[]int{-3, -3}, 3},
		{[]int{-1, 2}, 2},
		{[]int{-1, -1}, 1},
		{[]int{-2}, 2},
		{[]int{-1}, 1},
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{0, 0, 0}, 0},
		{[]int{1}, 1},
		{[]int{2}, 2},
		{[]int{1, 2}, 2},
		{[]int{3, 3}, 3},
		{[]int{2, 3, 4}, 12},
		{[]int{6, 10}, 30},
		{[]int{5, 10, 65}, 130},
		{[]int{75, 330, 225, 450}, 4950},
		{[]int{13, 89, 456}, 527592},
		{[]int{-2, -3, 4, 0}, 0},
		{[]int{2, 3, 4, 0}, 0},
		{[]int{1, 0}, 0},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			got, gotError := LCM(tC.input...)

			checkResult(t, tC.want, got, gotError)
		})
	}

	errorTestCases := []struct {
		desc      string
		input     []int
		wantError error
	}{
		{"|math.MinInt| is too large to store in an int", []int{math.MinInt}, ErrAbsoluteValueOfMinInt},
		{"The result is too large to store in an int", []int{math.MaxInt, 2}, ErrOverflowDetected},
	}

	for _, tC := range errorTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, gotError := LCM(tC.input...)

			checkError(t, gotError, tC.wantError)
		})
	}
}

func TestLCMBig(t *testing.T) {
	testCases := []struct {
		input []*big.Int
		want  *big.Int
	}{
		{[]*big.Int{big.NewInt(-2), big.NewInt(-3), big.NewInt(-4)}, big.NewInt(12)},
		{[]*big.Int{big.NewInt(-3), big.NewInt(-3)}, big.NewInt(3)},
		{[]*big.Int{big.NewInt(-1), big.NewInt(2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(-1), big.NewInt(-1)}, big.NewInt(1)},
		{[]*big.Int{big.NewInt(-2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(-1)}, big.NewInt(1)},
		{[]*big.Int{}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(1)}, big.NewInt(1)},
		{[]*big.Int{big.NewInt(2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(1), big.NewInt(2)}, big.NewInt(2)},
		{[]*big.Int{big.NewInt(3), big.NewInt(3)}, big.NewInt(3)},
		{[]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(4)}, big.NewInt(12)},
		{[]*big.Int{big.NewInt(6), big.NewInt(10)}, big.NewInt(30)},
		{[]*big.Int{big.NewInt(5), big.NewInt(10), big.NewInt(65)}, big.NewInt(130)},
		{[]*big.Int{big.NewInt(75), big.NewInt(330), big.NewInt(225), big.NewInt(450)}, big.NewInt(4950)},
		{[]*big.Int{big.NewInt(13), big.NewInt(89), big.NewInt(456)}, big.NewInt(527592)},
		{[]*big.Int{big.NewInt(-2), big.NewInt(-3), big.NewInt(4), big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(1), big.NewInt(0)}, big.NewInt(0)},
		{[]*big.Int{big.NewInt(math.MinInt), big.NewInt(2)}, new(big.Int).Abs(big.NewInt(math.MinInt))},
		{[]*big.Int{big.NewInt(math.MinInt), new(big.Int).Abs(big.NewInt(math.MinInt))}, new(big.Int).Abs(big.NewInt(math.MinInt))},
		{[]*big.Int{big.NewInt(math.MaxInt), big.NewInt(2)}, new(big.Int).Mul(big.NewInt(math.MaxInt), big.NewInt(2))},
	}

	for _, tC := range testCases {
		testName := fmt.Sprintf("Input: %v", tC.input)
		t.Run(testName, func(t *testing.T) {
			if got := LCMBig(tC.input...); got.Cmp(tC.want) != 0 {
				t.Errorf("Expected LCM: %v, actual LCM: %v", tC.want, got) // Can print the big.Int values OK.
			}
		})
	}
}

func ExampleGCD() {
	m, n := 48, 18
	gcd, err := GCD(m, n)
	if err != nil {
		fmt.Printf("Error calculating the GCD of %d and %d: %v", m, n, err)
	} else {
		fmt.Println("The GCD of", m, "and", n, "is", gcd)
	}

	// Output: The GCD of 48 and 18 is 6
}

func ExampleLCM() {
	m, n := 48, 18
	lcm, err := LCM(m, n)
	if err != nil {
		fmt.Printf("Error calculating the LCM of %d and %d: %v", m, n, err)
		return
	} else {
		fmt.Println("The LCM of", m, "and", n, "is", lcm)
	}

	// Output: The LCM of 48 and 18 is 144
}

func ExampleLCMBig() {
	m, n := new(big.Int), big.NewInt(3)
	m.SetString("10000000000000000000", 10)
	lcm := LCMBig(m, n)
	fmt.Println("The LCM of", m, "and", n, "is", lcm)

	// Output: The LCM of 10000000000000000000 and 3 is 30000000000000000000
}

func BenchmarkGCD(b *testing.B) {
	for b.Loop() {
		_, err := GCD(96, 48)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLCM(b *testing.B) {
	for b.Loop() {
		_, err := LCM(96, 48)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLCMBig(b *testing.B) {
	for b.Loop() {
		LCMBig(big.NewInt(96), big.NewInt(48))
	}
}
