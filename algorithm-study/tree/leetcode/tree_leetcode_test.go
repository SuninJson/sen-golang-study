package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}
	fmt.Println(isValidBST(root))
}

func TestIsCompleteTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 5}
	//root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 8}
	IsCompleteTree(root)
}

func TestPathTarget(t *testing.T) {
	root := buildTreeFromLevelOrder([]int{1, -2, -3, 1, 3, -2, 0, -1})
	target := pathTarget(root, -1)
	fmt.Println(target)
}
