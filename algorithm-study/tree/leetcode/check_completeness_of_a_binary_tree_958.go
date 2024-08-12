package leetcode

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/

func IsCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := make([]TreeNode, 1)
	queue[0] = *root

	nextMustBeLeaf := false

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if (cur.Left == nil && cur.Right != nil) || (nextMustBeLeaf && (cur.Left != nil || cur.Right != nil)) {
			return false
		}

		if cur.Left == nil || cur.Right == nil {
			nextMustBeLeaf = true
		}

		if cur.Left != nil {
			queue = append(queue, *cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, *cur.Right)
		}
	}

	return true
}
