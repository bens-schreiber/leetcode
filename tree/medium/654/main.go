package main

/*
PROMPT:
You are given an integer array with no duplicates. A binary tree can be
built by finding the maximum value in nums, and build the left subtree from all values to the left
of that index, and the right subtree as all values to the right of that index.

Solution:
Find the maximum value index via searching the list ( O(n) )
Create a prefix and suffix list from the max value index

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Return the index, and the value of the max int
func findMax(num []int) (int, int) {
	index, max := 0, num[0]
	for i, v := range num {
		if v > max {
			index, max = i, v
		}
	}
	return index, max
}

func makeTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex, maxValue := findMax(nums)

	root := &TreeNode{maxValue, nil, nil}

	if len(nums) == 1 {
		return root
	}

	prefix := nums[:maxIndex]
	suffix := nums[(maxIndex + 1):]

	root.Left = makeTree(prefix)
	root.Right = makeTree(suffix)
	return root
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return makeTree(nums)
}
