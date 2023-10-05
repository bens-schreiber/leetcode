/**

PROMPT:
You are given a binary tree such that the root is always equal to the minimum of its left and right nodes
Given such a tree, you need to output the second minimum value in the set of all the nodes

Intution:
A binary tree is a tree that has 0-2 nodes on every child

Questions:
What tests do I need to pass
What do I do if there is no second smallest value? (Return -1)

Solutions:
Min heapify the tree (nlogn)



**/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil || node.Right == nil
}

func findSecondMinimumValue(root *TreeNode) int {
	if isLeaf(root) {
		return -1
	}

	min := _findSecondMinimumValueImpl(root)

	// Min can never be the root. Error in this case.
	if min == root.Val {
		return -1
	}

	return min
}

func _findSecondMinimumValueImpl(node *TreeNode) int {
	if isLeaf(node.Left) || isLeaf(node.Right) {
		return max(node.Left.Val, node.Right.Val)
	}

	leftMin := _findSecondMinimumValueImpl(node.Left)
	rightMin := _findSecondMinimumValueImpl(node.Right)

	return min(leftMin, rightMin)
}

func main() {
	fmt.Println(findSecondMinimumValue(&TreeNode{2, &TreeNode{2, nil, nil}, &TreeNode{5, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}}}))
}
