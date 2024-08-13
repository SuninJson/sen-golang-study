package leetcode

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepth := 0

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		handTime := len(queue)

		for i := 0; i < handTime; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		maxDepth++
	}

	return maxDepth
}
