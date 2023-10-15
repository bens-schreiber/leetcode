package main

/*

PROMPT:
Given the root of a binary tree, sum all the values of nodes with even valued grandparents.
Return 0 if there are no even valued grandparents.
A grandparent is a parent of its parent.

Solution:
We can depth first search the tree, and everytime we get an even number, call a function
sumGrandChildren.

sumGrandChildren can be a breadth first search that goes two layers down, and returns the sum.


**** This algorithm is O(n^2), and is less efficient than other solutions ****

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumGrandChildren(root *TreeNode) int {
	queue := []*TreeNode{root}
	targetLayer := 3
	sum := 0
	for len(queue) > 0 && targetLayer > 0 {
		levelSize := len(queue)
		targetLayer--
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if targetLayer != 0 {
			sum = 0
		}
	}
	return sum
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	sum := left + right
	if root.Val%2 == 0 {
		sum += sumGrandChildren(root)
	}

	return sum
}

func sumEvenGrandparent(root *TreeNode) int {
	return dfs(root)
}
