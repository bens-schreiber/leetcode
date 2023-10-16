package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var dp = map[int][]*TreeNode{}

func backTrack(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	if n == 1 {
		return []*TreeNode{&TreeNode{Val: 0}}
	}

	// Check if it is in our cache
	val, ok := dp[n]
	if ok {
		return val
	}

	result := []*TreeNode{}
	for left := 0; left < n; left++ {
		right := n - 1 - left
		leftTrees, rightTrees := backTrack(left), backTrack(right)

		for _, t1 := range leftTrees {
			for _, t2 := range rightTrees {
				result = append(result, &TreeNode{0, t1, t2})
			}
		}
	}

	dp[n] = result
	return result
}

func allPossibleFBT(n int) []*TreeNode {
	if n % 2 == 0 {
		return []*TreeNode{}
	}
	return backTrack(n)
}
