package leetcode

// https://leetcode.cn/problems/house-robber-iii

type Result struct {
	robAmount    int
	notRobAmount int
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := doRob(root)

	return max(result.robAmount, result.notRobAmount)
}

func doRob(root *TreeNode) Result {
	if root == nil {
		return Result{0, 0}
	}

	leftResult := doRob(root.Left)
	rightResult := doRob(root.Right)

	robAmount := root.Val + leftResult.notRobAmount + rightResult.notRobAmount
	notRobAmount := max(leftResult.robAmount, leftResult.notRobAmount) + max(rightResult.robAmount, rightResult.notRobAmount)

	return Result{robAmount, notRobAmount}

}
