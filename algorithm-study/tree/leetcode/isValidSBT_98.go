package leetcode

import "fmt"

// https://leetcode.cn/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	inOrderArr := toInOrderArr(root)
	fmt.Println(inOrderArr)
	return isValidBSTByInOrderArr(inOrderArr)
}

func toInOrderArr(root *TreeNode) []int {
	inOrderArr := make([]int, 0)
	inOrderArr = getInOrderByStack(root)
	return inOrderArr
}

func inOrder(root *TreeNode, inOrderArr *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, inOrderArr)
	*inOrderArr = append(*inOrderArr, root.Val)
	inOrder(root.Right, inOrderArr)
}

func getInOrderByStack(cur *TreeNode) []int {
	inOrderArr := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		inOrderArr = append(inOrderArr, cur.Val)

		cur = cur.Right
	}
	return inOrderArr
}

func isValidBSTByInOrderArr(inOrderArr []int) bool {
	for i := 1; i < len(inOrderArr); i++ {
		if inOrderArr[i] < inOrderArr[i-1] {
			return false
		}
	}
	return true
}
