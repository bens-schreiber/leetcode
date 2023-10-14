package main

import "testing"

func TestNil(t *testing.T) {
	got := isBalanced(nil)
	want := true
	if got != want {
		t.Errorf("want %t got %t", want, got)
	}
}

func TestRoot(t *testing.T) {
	got := isBalanced(&TreeNode{1, nil, nil})
	want := true
	if got != want {
		t.Errorf("want %t got %t", want, got)
	}
}
