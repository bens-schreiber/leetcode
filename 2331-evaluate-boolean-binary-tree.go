/***

Prompt:
You are given a root of a full binary tree with the following properties
- Leaf nodes have either 0 or 1 indiciating false or true
- Non leaf nodes have either the value 2 or 3, where 2 represents OR and 3 represents AND

Intuition:
A binary tree is a tree that has nodes on the left and right
A full binary tree has either 2 children or 0 on all nodes
A leaf node has no children

Problems identified:
Search the tree with the given criteria
Know if a node is a leaf node or not

Questions:
What tests do I need to pass?
[2,1,3,null,null,0,1] -> true
[0] -> false

My solution:
Use a recursive depth first search


***/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func evaluateTree(root *TreeNode) bool {
	if isLeafNode(root) {
		return root.Val != 0
	}

	// OR
	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}

	return evaluateTree(root.Left) && evaluateTree(root.Right)
}

func main() {
	fmt.Println(evaluateTree(&TreeNode{0, nil, nil})) // false
	fmt.Println(evaluateTree(&TreeNode{1, nil, nil})) // true

	// // FALSE OR FALSE
	fmt.Println(evaluateTree(&TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}}))

	// // TRUE OR FALSE
	fmt.Println(evaluateTree(&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{0, nil, nil}}))

	// // TRUE OR TRUE
	fmt.Println(evaluateTree(&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{0, nil, nil}}))

	// // TRUE AND FALSE
	fmt.Println(evaluateTree(&TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{0, nil, nil}}))

	// // TRUE AND TRUE
	fmt.Println(evaluateTree(&TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}))
}
