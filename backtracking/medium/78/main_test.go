package main

import "testing"

func TestEmpty(t *testing.T) {
	got := subsets([]int{})
	want := [][]int{{}}

	if len(got) != len(want) {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestSingle(t *testing.T) {
	got := subsets([]int{1})
	want := [][]int{{1}, {}}

	if len(got) != len(want) {
		t.Fatalf("want %d got %d", want, got)
	}

	for i1, v1 := range want {
		if len(got[i1]) != len(v1) {
			t.Fatalf("want %d got %d", want, got)
		}
		for i2, v2 := range v1 {
			if got[i1][i2] != v2 {
				t.Fatalf("want %d got %d", want, got)
			}
		}
	}
}

// [1,2,3] -> [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func TestCase1(t *testing.T) {
	got := subsets([]int{1, 2, 3})
	want := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}

	if len(got) != len(want) {
		t.Fatalf("want %d got %d", want, got)
	}

	for i1, v1 := range want {
		if len(got[i1]) != len(v1) {
			t.Fatalf("want %d got %d", want, got)
		}
		for i2, v2 := range v1 {
			if got[i1][i2] != v2 {
				t.Fatalf("want %d got %d", want, got)
			}
		}
	}
}
