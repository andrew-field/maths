package maths

import (
	"sync"
)

// Inspiration taken from golang.org/x/tour/tree

// A Tree has a value and two sub trees.
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

	// Create all the tree nodes.
	for ind, val := range values {
		trees[ind] = &Tree{nil, val, nil}
	}

	pyramidSet, pyramidLimit := 0, 0

	// Take the tree nodes and link them together to build the tree structure.
	i, j := 0, 1
	for ; j < numberOfValues-1; i++ {
		trees[i].Left = trees[j]
		j++
		trees[i].Right = trees[j]
		switch {
		case !isPyramid:
			j++
		case pyramidSet == pyramidLimit:
			j++
			pyramidSet = 0
			pyramidLimit++
		default:
			pyramidSet++
		}
	}

	if j != numberOfValues {
		trees[i].Left = trees[j]
	}

	return trees[0]
}

// MaxPath returns the largest of all the possible summations from top to bottom of a tree.
// The execution works up from the bottom of the pyramid. The maximum path to a node is the value of the node
// plus the maximum of the maximum paths to each child node.
// There is a natural recursive function but it fails when a pyramid tree gets too large and the function runs out of resources.
// MaxPath(<nil>) returns 0.
func MaxPath(t *Tree) int {
	if t == nil {
		return 0
	}

	maximumPaths := new(sync.Map)

	for _, totalSumExists := maximumPaths.Load(t); !totalSumExists; _, totalSumExists = maximumPaths.Load(t) {
		generateMaximumPaths(t, maximumPaths, new(sync.Map))
	}

	sum, _ := maximumPaths.Load(t)
	return sum.(int)
}

func generateMaximumPaths(t *Tree, maximumPaths *sync.Map, channelsMap *sync.Map) {
	_, ok := channelsMap.LoadOrStore(t, true) // Prevents duplicate generating paths.
	if ok {
		return
	}

	var leftMax, rightMax interface{} = 0, 0
	leftMaxExists, rightMaxExists := true, true
	if t.Left != nil {
		leftMax, leftMaxExists = maximumPaths.Load(t.Left)
	}
	if !leftMaxExists {
		go generateMaximumPaths(t.Left, maximumPaths, channelsMap)
	}

	if t.Right != nil {
		rightMax, rightMaxExists = maximumPaths.Load(t.Right)
	}
	if !rightMaxExists {
		go generateMaximumPaths(t.Right, maximumPaths, channelsMap)
	}

	if leftMaxExists && rightMaxExists {
		maximumPaths.Store(t, Max(leftMax.(int), rightMax.(int))+t.Value)
	}
}
