package main

import "testing"

func TestRottingOranges(t *testing.T) {
	grid := [][]int{
		[]int{2, 1, 1},
		[]int{1, 1, 0},
		[]int{0, 1, 1},
	}
	expected := 4
	got := rottingOranges(grid)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
