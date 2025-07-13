package maths

import (
	"fmt"
	"slices"
	"strings"
)

// Inspiration taken from golang.org/x/tour/tree

// A Tree has a value and two sub trees.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) String() string {

	if t == nil {
		return ""
	}

	var b strings.Builder
	currentLevel := []*Tree{t}

	for {
		nextLevel := []*Tree{}
		for i, node := range currentLevel {
			if node == nil {
				return b.String()
			}
			if i > 0 {
				b.WriteString(" ")
			}
			fmt.Fprintf(&b, "%d", node.Value)
			if slices.Contains(nextLevel, node.Left) {
				nextLevel = append(nextLevel, node.Right)
			} else {
				nextLevel = append(nextLevel, node.Left, node.Right)
			}
		}
		if nextLevel[0] == nil {
			return b.String()
		}
		b.WriteString("\n")
		currentLevel = nextLevel
	}
}

// CreateBinaryTree returns a (mostly) symmetric binary tree, filling with values from top to bottom, left to right.
// Each row has double the number of nodes as the previous row, starting with 1 node at the top.
// The tree is not guaranteed to be complete, so the last row may not be full.
// CreateBinaryTree() returns <nil>
func CreateBinaryTree(values ...int) *Tree {
	return createTree(true, values...)
}

// CreatePyramidTree returns a (mostly) symmetric pyramid tree, filling with values from top to bottom, left to right.
// Each row has one more node than the previous row, starting with 1 node at the top.
// The tree is not guaranteed to be complete, so the last row may not be full.
// CreatePyramidTree() returns <nil>
func CreatePyramidTree(values ...int) *Tree {
	return createTree(false, values...)
}

func createTree(isBinaryTree bool, values ...int) *Tree {
	if len(values) == 0 {
		return nil
	}

	numberOfValues := len(values)

	trees := make([]*Tree, numberOfValues)

	// Create all the tree nodes, initialised with a value.
	for ind, val := range values {
		trees[ind] = &Tree{nil, val, nil}
	}

	i, j := 0, 1
	// Take the tree nodes and link them together to build the tree structure.
	if isBinaryTree {
		for ; j < numberOfValues-1; i++ {
			trees[i].Left = trees[j]
			j++
			trees[i].Right = trees[j]
			j++
		}
	} else {
		// pyramidLimit is the number of nodes that should have two parent nodes in the current row. This number increases by one for every row.
		// pyramidSet is the number of nodes that have two parent nodes in the current row. When pyramidSet equals pyramidLimit, the process moves to the next row.
		pyramidLimit, pyramidSet := 0, 0

		for ; j < numberOfValues-1; i++ {
			trees[i].Left = trees[j]
			j++
			trees[i].Right = trees[j]
			if pyramidSet == pyramidLimit {
				j++
				pyramidSet = 0
				pyramidLimit++
			} else {
				pyramidSet++
			}
		}
	}

	if j != numberOfValues {
		trees[i].Left = trees[j]
	}

	return trees[0]
}

// MaxPath returns the largest of all the possible summations from top to bottom of a tree.
// The execution works up from the bottom of the pyramid. The maximum path to a node is the value of the node plus the maximum of the maximum paths to each child node.
// MaxPath(<nil>) returns 0.
func MaxPath(t *Tree) int {
	if t == nil {
		return 0
	}

	maximumPaths := make(map[*Tree]int)
	maximumPaths[nil] = 0

	generateMaximumPaths(t, maximumPaths)

	return maximumPaths[t]
}

func generateMaximumPaths(t *Tree, maximumPaths map[*Tree]int) {
	_, ok := maximumPaths[t] // Prevents duplicate generating paths.
	if ok {
		return
	}

	if t.Left != nil {
		generateMaximumPaths(t.Left, maximumPaths)
	}

	if t.Right != nil {
		generateMaximumPaths(t.Right, maximumPaths)
	}

	maximumPaths[t] = max(maximumPaths[t.Left], maximumPaths[t.Right]) + t.Value
}
