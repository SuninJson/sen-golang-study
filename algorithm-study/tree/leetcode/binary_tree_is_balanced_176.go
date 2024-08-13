package leetcode

// https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/

// 二叉树的递归套路
//
//1）假设以X节点为头，假设可以向X左树和X右树要任何信息
//2）在上一步的假设下，讨论以X为头节点的树，得到答案的可能性（最重要）
//3）列出所有可能性后，确定到底需要向左树和右树要什么样的信息
//4）把左树信息和右树信息求全集，就是任何一棵子树都需要返回的信息S
//5）递归函数都返回S，每一棵子树都这么要求
//6）写代码，在代码中考虑如何把左树的信息和右树信息整合出整棵树的信息S

type IsBalancedRes struct {
	isBalanced bool
	height     int
}

func isBalanced(root *TreeNode) bool {
	return process(root).isBalanced
}

func process(cur *TreeNode) IsBalancedRes {
	if cur == nil {
		return IsBalancedRes{true, 0}
	}

	res := IsBalancedRes{}

	leftRes := process(cur.Left)
	rightRes := process(cur.Right)

	res.height = max(leftRes.height, rightRes.height) + 1

	if leftRes.isBalanced && rightRes.isBalanced && abs(leftRes.height-rightRes.height) <= 1 {
		res.isBalanced = true
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
