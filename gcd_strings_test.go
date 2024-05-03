package main

import "testing"

func TestGCDStrings(t *testing.T) {
	str1 := "ABCABC"
	str2 := "ABC"
	expected := "ABC"
	got := gcdOfStrings(str1, str2)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	str1 = "ABABAB"
	str2 = "ABAB"
	expected = "AB"
	got = gcdOfStrings(str1, str2)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	str1 = "LEET"
	str2 = "CODE"
	expected = ""
	got = gcdOfStrings(str1, str2)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
