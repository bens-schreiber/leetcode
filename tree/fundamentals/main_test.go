package main

import "testing"

func TestInOrderTraversal(t *testing.T) {
	got := inOrderTraversal(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}})
	want := []int{4, 2, 5, 1, 3}
	if len(got) != len(want) {
		t.Fatalf("got %d want %d", got, want)
		return
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %d want %d", got, want)
		}
	}
}

func TestPreOrderTraversal(t *testing.T) {
	got := preOrderTraversal(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}})
	want := []int{1, 2, 4, 5, 3}
	if len(got) != len(want) {
		t.Fatalf("got %d want %d", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %d want %d", got, want)
		}
	}
}

func TestPostOrderTraversal(t *testing.T) {
	got := postOrderTraversal(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}})
	want := []int{4, 5, 2, 3, 1}
	if len(got) != len(want) {
		t.Fatalf("got %d want %d", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %d want %d", got, want)
		}
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	got := levelOrderTraversal(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}})
	want := []int{1, 2, 3, 4, 5}
	if len(got) != len(want) {
		t.Fatalf("got %d want %d", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %d want %d", got, want)
		}
	}
}

func TestLevelOrderTraversalAware(t *testing.T) {
	got := levelOrderTraversalAware(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}})
	want := []int{1, 2, 3, 4, 5, 6, 7}
	if len(got) != len(want) {
		t.Fatalf("got %d want %d", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %d want %d", got, want)
		}
	}
}

func TestTreeHeight(t *testing.T) {
	got := treeHeight(&TreeNode{1, &TreeNode{1, &TreeNode{1, &TreeNode{1, nil, nil}, nil}, nil}, &TreeNode{1, nil, nil}})
	want := 3
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
