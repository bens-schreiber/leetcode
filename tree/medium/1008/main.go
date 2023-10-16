package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func peek(stack []*TreeNode) *TreeNode {
	return stack[len(stack)-1]
}

func pop(stack []*TreeNode) ([]*TreeNode, *TreeNode) {
	ret := peek(stack)
	stack = stack[:len(stack)-1]
	return stack, ret
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}

	for i := 1; i < len(preorder); i++ {
		node := &TreeNode{preorder[i], nil, nil}
		if node.Val < peek(stack).Val {
			peek(stack).Left = node
			stack = append(stack, node)
			continue
		}

		var parent *TreeNode
		for len(stack) > 0 && peek(stack).Val < node.Val {
			stack, parent = pop(stack)
		}

		parent.Right = node
		stack = append(stack, node)
	}

	return root
}
