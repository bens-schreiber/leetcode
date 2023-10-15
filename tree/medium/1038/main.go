package main

/*

PROMPT:
Given a binary search tree, convert it to a "Greater Search Tree" such that every node is the sum of the nodes greater than it.

Solution:
DFS the tree. Increment the node by the right node. A nil node has a value of 0.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, sum *int) {
	if node == nil {
		return
	}

	dfs(node.Right, sum)
	node.Val += *sum
	*sum = node.Val
	dfs(node.Left, sum)
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	dfs(root, &sum)
	return root
}

func main() {

}
