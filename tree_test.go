package maths

import "testing"

func TestMaxPath(t *testing.T) {
	testCases := []struct {
		input *Tree
		want  int
	}{
		{CreateBinaryTree(), 0},
		{CreateBinaryTree(1, 2, -3), 3},
		{CreateBinaryTree(-1), -1},
		{CreateBinaryTree(0), 0},
		{CreateBinaryTree(1), 1},
		{CreateBinaryTree(1, 2), 3},
		{CreateBinaryTree(1, 2, 2), 3},
		{CreateBinaryTree(1, 2, 3), 4},
		{CreateBinaryTree(1, 2, 3, 4, 5), 8},
		{CreateBinaryTree(75, 95, 64, 17, 47, 82, 18, 35, 87, 10), 274},
		{CreateBinaryTree(75, 95, 64, 17, 47, 82, 18, 35, 87, 10, 24, 12, 54, 38, 20, 46, 35, 42, 64, 21, 45), 338},
		{CreatePyramidTree(1, 2, -3, 4, -5, 6), 7},
		{CreatePyramidTree(1, 2, -3), 3},
		{CreatePyramidTree(-1), -1},
		{CreatePyramidTree(0), 0},
		{CreatePyramidTree(1), 1},
		{CreatePyramidTree(1, 2), 3},
		{CreatePyramidTree(1, 2, 2), 3},
		{CreatePyramidTree(1, 2, 3), 4},
		{CreatePyramidTree(1, 2, 3, 4, 5, 6), 10},
		{CreatePyramidTree(75, 95, 64, 17, 47, 82, 18, 35, 87, 10), 308},
		{CreatePyramidTree(75, 95, 64, 17, 47, 82, 18, 35, 87, 10, 24, 12, 54, 38, 20, 46, 35, 42, 64, 21, 45), 426},
	}

	for _, tC := range testCases {
		if got := MaxPath(tC.input); got != tC.want {
			t.Errorf("Input in test: %v. Actual result: %v. Expected result: %v", tC.input, got, tC.want)
		}
	}
}
