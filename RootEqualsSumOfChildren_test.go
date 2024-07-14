package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestCheckTree(t *testing.T) {
	root := TreeNode{
		Val:   10,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 6},
	}

	actual := checkTree(&root)

	assert.Equal(t, true, actual)
}

func checkTree(root *TreeNode) bool {
	return root.Left.Val+root.Right.Val == root.Val
}
