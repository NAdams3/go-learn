package main

import "testing"

func TestIntegerToBinary(t *testing.T) {

	expected := 1000
	got, _ := IntegerToBinary(8)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 1001
	got, _ = IntegerToBinary(9)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 110
	got, _ = IntegerToBinary(6)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 101000001
	got, _ = IntegerToBinary(321)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
