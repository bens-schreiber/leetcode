package main

/*
PROMPT:
Given a binary tree, return the number of nodes that are equal to the average of their subtree.

The average of a subtree is calculated by adding all of the nodes in the subtree, and dividing it by how many nodes are
in the tree.

Solution:
We can use a PostOrder traversal via DFS to solve this problem.

Create a DFS function that tracks the subtree sum, subtre node count, and our total averages found.
Calculate the left and right of those values.

If the averages equals the node that we are on, increase the averages.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// subtree sum, subtree node count, averages found
func dfs(node *TreeNode) (int, int, int) {
	if node == nil {
		return 0, 0, 0
	}

	leftSum, leftCount, leftAverages := dfs(node.Left)
	rightSum, rightCount, rightAverages := dfs(node.Right)

	sum := leftSum + rightSum + node.Val
	count := leftCount + rightCount + 1
	averages := leftAverages + rightAverages

	average := sum / count
	if node.Val == average {
		return sum, count, averages + 1
	}
	return sum, count, averages
}

func averageOfSubtree(root *TreeNode) int {
	_, _, average := dfs(root)
	return average

}
