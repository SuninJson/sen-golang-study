package leetcode

//https://leetcode.cn/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}
	queue := make([]TreeNode, 1)
	queue = append(queue, *root)

	for len(queue) > 0 {
		handleTime := len(queue)
		level := make([]int, 0)

		for i := 0; i < handleTime; i++ {
			cur := queue[0]
			queue = queue[1:]

			level = append(level, cur.Val)

			if cur.Left != nil {
				queue = append(queue, *cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, *cur.Right)
			}
		}
	}

	return result
}
