package main

/*

PROMPT:
Given the root of a binary tree, return the sums of the values of its deepest leaves.

Solution:
BFS the tree. Save a copy of the queue of nodes. When the BFS is done, the copy of the nodes should be the last layer of nodes.
Sum that layer.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	lastLayer := []*TreeNode{}
	for len(queue) > 0 {
		levelSize := len(queue)
		lastLayer = queue
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	sum := 0
	for _, v := range lastLayer {
		sum += v.Val
	}
	return sum
}

func main() {

}
