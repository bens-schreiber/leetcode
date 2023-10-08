package main

import (
	"testing"
)

/*

PROMPT:
Find the minimum path from the root of a binary tree to a no children

Intution:
A binary tree is a tree such that there are 0-2 children on all nodes

Questions:
What tests should I be able to pass?

[3,9,20,null,null,15,7] -> 2
[2,null,3,null,4,null,5,null,6] -> 5

Solutions:
We should use a breadth first search. Meaning, we go layer by layer of the tree, and check for both children being null.

In psuedo code:
int layer = 0
for each layer in tree:
 for each child in layer:
	if child has no children:
		return layer
 layer++


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodes := []*TreeNode{root}
	layer := 0

	for len(nodes) != 0 {
		layer++
		nodeCount := len(nodes)

		for i := 0; i < nodeCount; i++ {
			node := nodes[0]
			nodes = nodes[1:]

			if node.Left == nil && node.Right == nil {
				return layer
			}

			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}

			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
	}

	return layer
}

func TestSingletonTree(t *testing.T) {
	got := minDepth(&TreeNode{0, nil, nil})
	want := 1
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// [3,9,20,null,null,15,7] -> 2
func Test1(t *testing.T) {
	got := minDepth(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}})
	want := 2
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// [1,2,3,4,5] -> 2
func Test2(t *testing.T) {
	got := minDepth(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, &TreeNode{5, nil, nil}, nil}})
	want := 2
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func main() {

}
