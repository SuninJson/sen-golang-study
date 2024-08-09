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
