package main

import "testing"

func TestRoot(t *testing.T) {
	got := isSymmetric(&TreeNode{1, nil, nil})
	want := true
	if got != want {
		t.Fatalf("got %t want %t", got, want)
	}
}
