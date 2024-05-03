package main

import "testing"

func TestStringCompression(t *testing.T) {
	chars := []string{"a", "a", "b", "b", "c", "c", "c"}
	expected := 6
	got := compressChars(chars)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
