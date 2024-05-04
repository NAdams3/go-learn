package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	expected := true
	got := canPlaceFlowers(flowerbed, n)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	n = 2
	expected = false
	got = canPlaceFlowers(flowerbed, n)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	flowerbed = []int{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0}
	n = 3
	expected = false
	got = canPlaceFlowers(flowerbed, n)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	flowerbed = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	n = 6
	expected = true
	got = canPlaceFlowers(flowerbed, n)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
