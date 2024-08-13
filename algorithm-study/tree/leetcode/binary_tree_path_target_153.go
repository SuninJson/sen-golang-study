package leetcode

import "fmt"

type PathNode struct {
	Node *TreeNode
	Path []int
	Sum  int
}

func pathTarget(root *TreeNode, target int) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	queue := []PathNode{{Node: root, Path: []int{root.Val}, Sum: root.Val}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		fmt.Println(cur)

		curNode := cur.Node
		curPath := cur.Path
		curSum := cur.Sum

		if curNode.Left == nil && curNode.Right == nil && curSum == target {
			fmt.Println(curPath)
			result = append(result, append([]int{}, curPath...))
		}

		left := curNode.Left
		fmt.Println("left:", left)
		if left != nil {
			leftPathSum := curSum + left.Val
			fmt.Printf("leftPathSum := curSum + right.Val | leftPathSum:%d curSum: %d right.Val: %d  \n", leftPathSum, curSum, left.Val)

			leftPath := append(append([]int{}, curPath...), left.Val)
			leftPathNode := PathNode{
				Node: left,
				Path: leftPath,
				Sum:  leftPathSum,
			}
			fmt.Printf("curPath:%v  left.Val:%d leftPathNode:%v \n", curPath, left.Val, leftPathNode)
			queue = append(queue, leftPathNode)

		}

		fmt.Println("handle left queue:", queue)

		fmt.Println("right:", curNode.Right)
		if curNode.Right != nil {
			rightPathSum := curSum + curNode.Right.Val
			fmt.Printf("rightPathSum := curSum + right.Val | rightPathSum: %d curSum: %d  right.Val:%d \n", rightPathSum, curSum, curNode.Right.Val)

			rightPath := append(append([]int{}, curPath...), curNode.Right.Val)
			rightPathNode := PathNode{
				Node: curNode.Right,
				Path: rightPath,
				Sum:  rightPathSum,
			}
			fmt.Printf("curPath:%v  right.Val:%d rightPathNode:%v \n", curPath, curNode.Right.Val, rightPathNode)
			queue = append(queue, rightPathNode)

		}

		fmt.Println("handle right queue:", queue)

		fmt.Println()
	}

	return result

}
