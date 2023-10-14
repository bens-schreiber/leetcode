package main

import "testing"

func TestNil(t *testing.T) {
	got := isCousins(nil, 0, 0)
	want := false
	if got != want {
		t.Fatalf("got %t want %t", got, want)
	}
}

func TestRoot(t *testing.T) {
	got := isCousins(&TreeNode{0, nil, nil}, 0, 0)
	want := false
	if got != want {
		t.Fatalf("got %t want %t", got, want)
	}
}

func TestRootWithChildren(t *testing.T) {
	got := isCousins(&TreeNode{0, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}, 1, 2)
	want := false
	if got != want {
		t.Fatalf("got %t want %t", got, want)
	}
}

func TestCase1(t *testing.T) {
	got := isCousins(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, &TreeNode{5, nil, nil}}}, 4, 5)
	want := true
	if got != want {
		t.Fatalf("got %t want %t", got, want)
	}
}

// 1, 2, 3, nil, nil, 4, 5
func TestCase2(t *testing.T) {
	got := isCousins(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 4, 5)
	want := true
	if got != want {
		t.Fatalf("got %t want %t", got, want)
	}
}
