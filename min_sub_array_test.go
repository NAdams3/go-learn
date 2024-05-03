package main

import "testing"

func TestMinSubArray(t *testing.T) {
	arr := []uint{4, 3, 7, 9}
	k := 4
	got := minimumSubArray(arr, k)
	expected := 2
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	arr = []uint{1, 2, 4, 8}
	k = 2
	got = minimumSubArray(arr, k)
	expected = 2
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
