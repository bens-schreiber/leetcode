/**

PROMPT:
Given an original tree, a copy of the original tree, and a target, find the reference to the
target in the original tree. (binary tree)

Intuition:
A binary tree is a tree that has 0-2 children

Questions:
What tests do I need to pass?

Solution:
Write a depth first search to find the target node. The search can be recursive,
and return either a node pointer or nil.

**/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findClone(original *TreeNode, copy *TreeNode, target int) *TreeNode {

	if original == nil || copy == nil {
		return nil
	}

	if original.Val == target && copy.Val == target {
		return copy
	}

	clone := findClone(original.Left, copy.Left, target)
	if clone != nil {
		return clone
	}

	clone = findClone(original.Right, copy.Right, target)
	return clone
}

func main() {

	// nil
	fmt.Println(findClone(nil, nil, 0))

	// nil
	fmt.Println(findClone(&TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}, 1))

	// ptr
	fmt.Println(findClone(&TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}, 0))
}
