package main

/*
PROMPT:
Find the longest path in a binary tree.

Intuition:
The longest path is going to be from the deepest left node to the deepest right node. We need to find those two nodes.

Solution:
DFS the right and left. The longest path is going to be the largest sum of the right and left depths.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, max := dfs(root)
	return max
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftDepth, leftMax := dfs(root.Left)
	rightDepth, rightMax := dfs(root.Right)

	depth := max(leftDepth, rightDepth) + 1
	max := max(leftDepth+rightDepth, max(leftMax, rightMax))

	return depth, max
}
