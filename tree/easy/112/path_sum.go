package main

/*

PROMPT:
Given a binary tree root, and some integer target sum, return true if there is a path from the root
that adds to the target sum.

Intuition:
A binary tree is a tree such that there are 0-2 nodes for each child. This tree is not gaurunteed full nor is it
balanced, so it could be skewed. Our path must be root to leaf.

Since we are looking for a path from the root, we have at best O(n) time for this algorithm.

Questions:
What unit tests do I need to pass?
Do my trees always contain a target?
Are the values always positive?

Solution:
We will do a Depth First Search.
If the node is nil, return false.
If the value is equal to the target sum, and the value is a leaf, return true.
Else, go down the left path and the right path with the same logic, but decrementing the targetSum by the root.Val

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(node *TreeNode) bool {
	return node == nil || node.Left == nil && node.Right == nil
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && isLeaf(root) {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {

}
