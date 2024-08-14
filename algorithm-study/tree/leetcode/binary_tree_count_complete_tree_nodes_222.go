package leetcode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return doCount(root, 1, mostLeftDepth(root, 1))
}

func mostLeftDepth(root *TreeNode, height int) int {
	cur := root
	for cur != nil {
		height++
		cur = cur.Left
	}

	return height - 1
}

func doCount(root *TreeNode, levelNo int, wholeTreeHeight int) int {
	if levelNo == wholeTreeHeight {
		// 若当前节点已处于到最后一层，则以 当前节点为根节点的树 的节点数为 1
		return 1
	}

	rightMostLeftDepth := mostLeftDepth(root.Right, levelNo+1)
	if rightMostLeftDepth == wholeTreeHeight {
		// 若当前节点的右树的最左节点处于最底层，则当前节点的左树为满二叉树
		return (1 << (wholeTreeHeight - levelNo)) + doCount(root.Right, levelNo+1, wholeTreeHeight)
	} else {
		// 若当前节点的右树的最左节点不处于最底层，则当前节点的右树为满二叉树，但其层数要比左树少 1 层
		return (1 << (wholeTreeHeight - levelNo - 1)) + doCount(root.Left, levelNo+1, wholeTreeHeight)
	}
}
