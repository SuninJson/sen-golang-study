package leetcode

func buildTreeFromLevelOrder(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		leftVal := nums[i]
		i++
		if leftVal != 0 { // -1 表示空节点
			leftNode := &TreeNode{Val: leftVal}
			node.Left = leftNode
			queue = append(queue, leftNode)
		}

		if i < len(nums) {
			rightVal := nums[i]
			i++
			if rightVal != 0 { // -1 表示空节点
				rightNode := &TreeNode{Val: rightVal}
				node.Right = rightNode
				queue = append(queue, rightNode)
			}
		}
	}

	return root
}
