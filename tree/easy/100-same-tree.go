/*

PROMPT:
Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

EX 1:
Input: p = [1,2,3], q = [1,2,3]
Output: true

EX 2:
Input: p = [1,2], q = [1,null,2]
Output: false

EX 3:
Input: p = [1,2,1], q = [1,1,2]
Output: false

Intuition:
A binary tree is a tree datastructure that has at most two children on any given node.
To check if the binary tree is the same, we need to compare the values of the nodes and the structure of the tree.

Questions:
How many nodes are in the tree?
What tests does this need to pass?

Solution:
A depth first search that compares the values of the nodes at every level.

*/

package main
import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

    if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	
}

func main() {
	fmt.Println(isSameTree(nil, nil))
	fmt.Println(isSameTree(&TreeNode{1,nil,nil}, &TreeNode{1,nil,nil}))
	fmt.Println(isSameTree(&TreeNode{1,nil,nil}, &TreeNode{2,nil,nil}))
}