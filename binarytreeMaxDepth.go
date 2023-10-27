package main

type TreeNode struct {
	   Val int
	   Left *TreeNode
	   Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}


	return 1 + intMax(maxDepth(root.Left), maxDepth(root.Right))
}

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}



