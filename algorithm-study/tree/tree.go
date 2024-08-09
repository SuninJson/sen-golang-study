package tree

import (
	"fmt"
)

type Node interface {
	PrintRecursiveOrder()
	GetVal() any
	GetLeft() Node
	GetRight() Node
}

type IntNode struct {
	Val       int
	LeftNode  Node
	RightNode Node
}

func (node *IntNode) GetRight() Node {
	return node.RightNode
}

func (node *IntNode) GetLeft() Node {
	return node.LeftNode
}

func (node *IntNode) PrintRecursiveOrder() {
	RecursiveOrderHandleTree(node, printNodeVal)
}

func (node *IntNode) GetVal() any {
	return node.Val
}

// RecursiveOrderHandleTree 按递归顺序处理树
// 注意每个结点都会被处理3次
func RecursiveOrderHandleTree(node Node, handeFunc func(node Node)) {
	if node == nil {
		return
	}

	handeFunc(node)
	RecursiveOrderHandleTree(node.GetLeft(), handeFunc)

	handeFunc(node)
	RecursiveOrderHandleTree(node.GetRight(), handeFunc)

	handeFunc(node)
}

func printNodeVal(node Node) {
	fmt.Print(" ", node.GetVal())
}
