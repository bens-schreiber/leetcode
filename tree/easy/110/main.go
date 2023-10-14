package main

import "math"

/*

PROMPT: Given a binary tree, determine if it is height balanced.

Height balanced - The depth of every nodes subtree does not
differ by more than 1

Intuition:

A binary tree is a tree such that there are 0-2 children

The depth of a binary tree is its path from the root to a leaf

Questions:

What tests do I need to pass?

Solution:
Calculate the tree height recursively.
If the left height or right height has an absolute difference greater than 1,
return false


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := treeHeight(node.Left)
	if left == -1 {
		return -1
	}

	right := treeHeight(node.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left-right)) > 1 {
		return -1
	}

	return max(left, right) + 1

}

func isBalanced(node *TreeNode) bool {
	return treeHeight(node) != -1
}

// O(n^2)
//
// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	left := treeDepth(root.Left)
// 	right := treeDepth(root.Right)
// 	if math.Abs(float64(left-right)) > 1 {
// 		return false
// 	}
// 	return isBalanced(root.Left) && isBalanced(root.Right)
// }

// func treeDepth(node *TreeNode) int {
// 	if node == nil {
// 		return 0
// 	}

// 	left := treeDepth(node.Left)
// 	right := treeDepth(node.Right)
// 	if left > right {
// 		return left + 1
// 	}
// 	return right + 1
// }

func main() {

}
