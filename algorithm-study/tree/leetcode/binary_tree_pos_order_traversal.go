package leetcode

// https://leetcode.cn/problems/binary-tree-postorder-traversal/

func postorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	tempStack := make([]*TreeNode, 1)
	tempStack[0] = root
	posOrderStack := make([]int, 0)

	for len(tempStack) > 0 {
		cur := tempStack[len(tempStack)-1]
		tempStack = tempStack[:len(tempStack)-1]

		posOrderStack = append(posOrderStack, cur.Val)

		if cur.Left != nil {
			tempStack = append(tempStack, cur.Left)
		}

		if cur.Right != nil {
			tempStack = append(tempStack, cur.Right)
		}
	}
	elementNum := len(posOrderStack)
	result := make([]int, elementNum)
	for i := 0; i < elementNum; i++ {
		result[i] = posOrderStack[elementNum-i-1]
	}

	return result
}
