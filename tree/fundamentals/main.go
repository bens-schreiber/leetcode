package main

// Basic tree definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
The in order traversal of a tree can be described as a depth first search, printing
the left, root, and right and then unwinding the recursion to the parent.
*/
func inOrderTraversal(root *TreeNode) []int {
	ret := []int{}

	if root == nil {
		return ret
	}

	// Left of the tree
	ret = append(ret, inOrderTraversal(root.Left)...)

	// Root of the tree
	ret = append(ret, root.Val)

	// Right of the tree
	ret = append(ret, inOrderTraversal(root.Right)...)

	return ret
}

/*
A pre-order traversal can be described as a depth first search yielding root, left, right
*/
func preOrderTraversal(root *TreeNode) []int {
	ret := []int{}

	if root == nil {
		return ret
	}

	// Root of the tree
	ret = append(ret, root.Val)

	// Left of the tree
	ret = append(ret, preOrderTraversal(root.Left)...)

	// Right of the tree
	ret = append(ret, preOrderTraversal(root.Right)...)

	return ret
}

/*
A post-order traversal can be described as a depth first search yielding left, right, root
*/
func postOrderTraversal(root *TreeNode) []int {
	ret := []int{}

	if root == nil {
		return ret
	}

	// Left of the tree
	ret = append(ret, postOrderTraversal(root.Left)...)

	// Right of the tree
	ret = append(ret, postOrderTraversal(root.Right)...)

	// Root of the tree
	ret = append(ret, root.Val)

	return ret
}

/*
A level order traversal can be described as a breadth first search:

for layer in tree:

	for child in layer:
		print child
*/
func levelOrderTraversal(root *TreeNode) []int {
	ret := []int{}
	q := []*TreeNode{root}
	var node *TreeNode = nil
	for len(q) > 0 {
		node = q[0]
		ret = append(ret, node.Val)
		q = q[1:]

		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return ret

}

/*
The tree height is described as the longest path from the root to a leaf.
We can calculate this with a DFS, comparing the length of the right and left, and returning the highest.
*/
func treeHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}

	left := 1 + treeHeight(root.Left)
	right := 1 + treeHeight((root.Right))

	if left > right {
		return left
	}

	return right

}

func main() {

}
