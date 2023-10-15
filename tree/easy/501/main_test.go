package main

import "testing"

func TestNil(t *testing.T) {
	got := findMode(nil)
	want := []int{}

	if len(got) != len(want) {
		t.Errorf("want %d got %d", want, got)
		return
	}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("want %d got %d", want, got)
		}
	}
}

func TestSingle(t *testing.T) {
	got := findMode(&TreeNode{1, nil, nil})
	want := []int{1}

	if len(got) != len(want) {
		t.Errorf("want %d got %d", want, got)
		return
	}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("want %d got %d", want, got)
		}
	}
}

func TestCase1(t *testing.T) {
	got := findMode(&TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}})
	want := []int{1}

	if len(got) != len(want) {
		t.Errorf("want %d got %d", want, got)
		return
	}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("want %d got %d", want, got)
		}
	}
}

func TestCase2(t *testing.T) {
	got := findMode(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}})
	want := []int{1, 2, 3}

	if len(got) != len(want) {
		t.Errorf("want %d got %d", want, got)
		return
	}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("want %d got %d", want, got)
		}
	}
}
