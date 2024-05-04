package main

import "testing"

func TestIncreasingTriplet(t *testing.T) {

	arr := []int{20, 100, 10, 12, 5, 13}
	expected := true
	got := increasingTriplet(arr)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	arr = []int{1, 2, 3, 4, 5}
	expected = true
	got = increasingTriplet(arr)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	arr = []int{5, 4, 3, 2, 1}
	expected = false
	got = increasingTriplet(arr)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
