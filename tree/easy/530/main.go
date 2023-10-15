package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt64
	prev := -1

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		if prev != -1 && node.Val-prev < minDiff {
			minDiff = node.Val - prev
		}
		prev = node.Val
		traverse(node.Right)
	}

	traverse(root)
	return minDiff
}
