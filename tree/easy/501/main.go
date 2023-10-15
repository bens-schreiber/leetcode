package main

/*

PROMPT: Given a binary search tree with duplicates,
return all the modes.

Modes - the most frequently occured element in the tree

Intuition:
A binary search tree is a binary tree such that there are 0-2
children on each node.

Each node of a BST has a larger value node on the right, and a smaller on the left (or nil)

We are looking at an O(n) best time algorithm because we are required to iterate over all nodes of the tree.

Questions:

Q: A binary search tree cannot have duplicates. Where does the equal node go?
A: The left or right. In this BST, its greater than or equal to instead of strictly greater than

Q: If we get equivalent modes, do I return both as a list?
A: Yes

Q: What should a root with no children return?
A: The root

What test cases should I pass?

Solution:
We know that we have to do a search of this tree, so I will go for a depth first search,
since it is the least space complexity.

To keep count of each of the nodes and their subsequent mode we should use a map
The (K,V) of the map is the TreeNode.Val, and the count of the node duplicates

Depth first search the tree, add the values to the map, and then increment the count.
Once the map is filled, iterate over all the values of the map.


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	m := make(map[int]int)
	_findModeImpl(root, m)

	highestDuplicate := 0
	l := []int{}
	for key, value := range m {
		if value > highestDuplicate {
			l = []int{key}
			highestDuplicate = value
			continue
		}
		if value == highestDuplicate {
			l = append(l, key)
		}
	}
	return l
}

func _findModeImpl(node *TreeNode, m map[int]int) {
	if node == nil {
		return
	}

	m[node.Val]++

	_findModeImpl(node.Left, m)
	_findModeImpl(node.Right, m)
}

func main() {

}
