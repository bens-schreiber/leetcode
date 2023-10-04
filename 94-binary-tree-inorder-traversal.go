/** 

PROMPT: Given the root of a binary tree, return the inorder traversal of its nodes' values.

EX 1:
Input: root = [1,null,2,3]
Output: [1,3,2]

EX 2:
Input: root = [1]
Output: [1]

Intuition:
A binary tree is a tree datastructure that has at most two children on any given node.
In order traversal is the act of traversing the tree in the order of left, root, right.

Questions:
How many nodes are in the tree?
What tests does this need to pass?
**/


package main
import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func inorderTraversal(root *TreeNode) []int {
	traversal := make([]int, 0)
	traversal = _impl(root, traversal)
	return traversal
}

func _impl(node *TreeNode, traversal []int) []int {
	if (node == nil) {
		return traversal
	}

	traversal = _impl(node.Left, traversal)
	traversal = append(traversal, node.Val)
	traversal = _impl(node.Right, traversal)
	return traversal
}

func main() {

	fmt.Println(inorderTraversal(nil))
	
	fmt.Println(inorderTraversal(&TreeNode{1, nil, nil}))

	fmt.Println(inorderTraversal(&TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, &TreeNode{5, nil, nil}}, nil}, nil}, nil}))
	
}