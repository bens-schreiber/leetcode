package main

import (
	"testing"
)

func TestNoChildren(t *testing.T) {
	got := findSecondMinimumValue(&TreeNode{0, nil, nil})
	want := -1
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestEquivalentChildren(t *testing.T) {
	got := findSecondMinimumValue(&TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}})
	want := -1
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCase1(t *testing.T) {
	got := findSecondMinimumValue(&TreeNode{2, &TreeNode{2, nil, nil}, &TreeNode{5, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}}})
	want := 5
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// [1,1,2,1,1,2,2]
func TestCase2(t *testing.T) {
	got := findSecondMinimumValue(&TreeNode{1, &TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{2, &TreeNode{2, nil, nil}, &TreeNode{2, nil, nil}}})
	want := 2
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
