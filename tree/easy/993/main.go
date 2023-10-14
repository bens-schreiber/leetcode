package main

/*
PROMPT:
Given the root of a binary tree with unique values, and the values of two nodes
within the tree x and y, return true if x and y are cousins.

x and y are cousins IFF they have the same depth, with different parents

Questions:

Q: What is the depth of a tree with no children?
A: 0

Solution:
Use a DFS to find x, and x's parent, and depth
Use a DFS to find y, and y's parent, and depth
If x's parent != y's parent, and x's depth == y's depth, return true

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	xLength := -1
	xPre := root
	yLength := -1
	yPre := root

	var findDepth func(*TreeNode, *TreeNode, int)
	findDepth = func(root *TreeNode, pre *TreeNode, depth int) {
		if root == nil {
			return
		}

		if root.Val == x {
			xLength = depth
			xPre = pre
		}

		if root.Val == y {
			yLength = depth
			yPre = pre
		}

		findDepth(root.Left, root, depth+1)
		findDepth(root.Right, root, depth+1)
	}

	findDepth(root, nil, 0)

	return xLength == yLength && xPre != yPre
}
