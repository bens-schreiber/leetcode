package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(root *TreeNode) bool {
	return root != nil && root.Left == nil && root.Right == nil
}
func sumOfLeftLeaves(root *TreeNode) int {
	if isLeaf(root) {
		return 0
	}
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if isLeaf(root) {
		return root.Val
	}

	left := dfs(root.Left)
	right := 0
	if !isLeaf(root.Right) {
		right += dfs(root.Right)
	}

	return left + right
}
