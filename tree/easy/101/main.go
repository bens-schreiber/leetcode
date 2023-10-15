package main

/*
PROMPT: Given the root of a binary tree, return true
if it is a mirror of itself (left branch mirrors right branch)

Intuition:
For the tree to be a mirror of itself, the left and right trees must be equal,
or be exactly opposite of each other (ie a left child on the left tree should have a right child on the right tree, and vice versa)

Solution:
We can determine the solution by using a depth first search.
If both the first node is nil and the second node is nil, return true
If one of them is nil, return false
If the values are not equal, return false
Recursively call the function with the left of node1 and the right of node2, and the right of node1 with the left of node2

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}
