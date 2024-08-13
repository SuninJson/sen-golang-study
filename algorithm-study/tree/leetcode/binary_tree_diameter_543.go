package leetcode

// https://leetcode.cn/problems/diameter-of-binary-tree/

type DiameterInfo struct {
	height int
	// 结果需要是边数，为了方便处理我们求直径上的节点数
	numOfNodesOnDiameter int
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 直径为节点之间的边数，节点之间的边数为节点数量减一，例如4个节点之间有3条边
	return getDiameter(root).numOfNodesOnDiameter - 1
}

func getDiameter(root *TreeNode) DiameterInfo {
	if root == nil {
		return DiameterInfo{0, 0}
	}

	leftInfo := getDiameter(root.Left)
	rightInfo := getDiameter(root.Right)

	height := max(leftInfo.height, rightInfo.height) + 1

	// 若直径上有根节点，则直径上的节点数为 左树的高度 + 1 + 右树的高度
	includeRoot := leftInfo.height + rightInfo.height + 1

	// 若直径上无根节点，则直径上的节点数为 左树的直径上节点数和右树上的节点树的较大值
	notIncludeRoot := max(leftInfo.numOfNodesOnDiameter, rightInfo.numOfNodesOnDiameter)

	return DiameterInfo{height, max(includeRoot, notIncludeRoot)}
}
