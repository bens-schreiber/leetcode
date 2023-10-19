/*
PROMPT:
Given an array, return all possible unique combinations of the numbers in that array (the power set)

Intuition:
We will need to generate all permutations of a set.

Questions:

Q: Max length of the array?
A: -10 <= nums[i] <= 10

Q: What unit tests do I need to pass?
A: [1,2,3] -> [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]


*/

package main

func subsets(nums []int) [][]int {
	result := [][]int{}
	subset := []int{}

	var dfs func(int)
	dfs = func(i int) {
		if i >= len(nums) {
			temp := make([]int, len(nums))
			copy(temp, subset)
			result = append(result, temp)
			return
		}

		// include
		subset = append(subset, nums[i])
		dfs(i + 1)

		// do nothing, unnapend
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)

	return result
}

func main() {

}
