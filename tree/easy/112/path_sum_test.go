package main

import "testing"

func TestNoChildren(t *testing.T) {
	got := hasPathSum(nil, 0)
	want := false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestOnlyRootFalse(t *testing.T) {
	got := hasPathSum(&TreeNode{0, nil, nil}, 1)
	want := false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestOnlyRootTrue(t *testing.T) {
	got := hasPathSum(&TreeNode{1, nil, nil}, 1)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestBasicTrue(t *testing.T) {
	got := hasPathSum(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}}, 3)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestBasicFalse(t *testing.T) {
	got := hasPathSum(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}}, 4)
	want := false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

// [-2,null,-3]
func TestTrue(t *testing.T) {
	got := hasPathSum(&TreeNode{-2, nil, &TreeNode{-3, nil, nil}}, -5)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
