package maths

// Inspiration taken from golang.org/x/tour/tree

// A Tree is has a value and two sub trees.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// CreateBinaryTree returns a (mostly) symmetric binary tree, filling with values from top to bottom, left to right.
// CreateBinaryTree() returns <nil>
func CreateBinaryTree(values ...int) *Tree {
	return createTree(false, values...)
}

// CreatePyramidTree returns a (mostly) symmetric pyramid tree, filling with values from top to bottom, left to right.
// CreatePyramidTree() returns <nil>
func CreatePyramidTree(values ...int) *Tree {
	return createTree(true, values...)
}

func createTree(isPyramid bool, values ...int) *Tree {
	if len(values) == 0 {
		return nil
	}

	numberOfValues := len(values)

	trees := make([]*Tree, numberOfValues)

	for ind, val := range values {
		trees[ind] = &Tree{nil, val, nil}
	}

	pyramidSet, pyramidLimit := 0, 0

	i, j := 0, 1
	for ; j < numberOfValues-1; i++ {
		trees[i].Left = trees[j]
		j++
		trees[i].Right = trees[j]
		if !isPyramid {
			j++
		} else if pyramidSet == pyramidLimit {
			j++
			pyramidSet = 0
			pyramidLimit++
		} else {
			pyramidSet++
		}
	}

	if j != numberOfValues {
		trees[i].Left = trees[j]
	}

	return trees[0]
}

// MaxPath returns the largest of all the possible summations from top to bottom of a tree.
// MaxPath(<nil>) returns 0
func MaxPath(t *Tree) int {
	if t == nil {
		return 0
	}

	return Max(MaxPath(t.Left), MaxPath(t.Right)) + t.Value
}
