package main

import "testing"

func TestLinkedList(t *testing.T) {
	list := LinkedList[int]{}

	list.append(5)
	list.append(7)
	list.append(9)

	got, _ := list.get(2)
	if got != 9 {
		t.Errorf("expected %d, got %d", 9, got)
	}

	got, _ = list.removeAt(1)
	if got != 7 {
		t.Errorf("expected %d, got %d", 7, got)
	}

	got = list.length
	if got != 2 {
		t.Errorf("expected %d, got %d", 2, got)
	}

	list.append(11)

	expected := 9
	got, _ = list.removeAt(1)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 5
	got, _ = list.removeAt(0)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 11
	got, _ = list.removeAt(0)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 0
	got = list.length
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	list.prepend(5)
	list.prepend(7)
	list.prepend(9)

	expected = 5
	got, _ = list.get(2)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 9
	got, _ = list.get(0)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 9
	got, _ = list.remove(9)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 2
	got = list.length
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 7
	got, _ = list.get(0)
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	_, err := list.get(2)
	if err == nil {
		t.Error("expected error, got nil")
	}

	_, err = list.removeAt(2)
	if err == nil {
		t.Error("expected error, got nil")
	}

	_, err = list.remove(22)
	if err == nil {
		t.Error("expected error, got nil")
	}

}
