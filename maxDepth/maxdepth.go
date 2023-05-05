package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode // Children
	Right *TreeNode // Children
}

// maxDepth
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int = 1
	var left int
	var right int
	if root.Left != nil {
		left = maxDepth(root.Left)
	}
	if root.Right != nil {
		right = maxDepth(root.Right)
	}

	m := max(left, right)

	return depth + m
}

// max
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}


func main() {
	tn := TreeNode{3, &TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	md := maxDepth(&tn)
	fmt.Println(md)
}
