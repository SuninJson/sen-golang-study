package tree

import (
	"testing"
)

func getIntNode() *IntNode {
	return &IntNode{
		Val: 1,
		LeftNode: &IntNode{
			Val: 11,
			LeftNode: &IntNode{
				Val:       111,
				LeftNode:  nil,
				RightNode: nil,
			},
			RightNode: &IntNode{
				Val:       112,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
		RightNode: &IntNode{
			Val: 12,
			LeftNode: &IntNode{
				Val:       121,
				LeftNode:  nil,
				RightNode: nil,
			},
			RightNode: &IntNode{
				Val:       122,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
	}
}

func TestIntNodePrintRecursiveOrder(t *testing.T) {
	intNode := getIntNode()
	intNode.PrintRecursiveOrder()
}
