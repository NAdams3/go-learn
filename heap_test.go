package main

import "testing"

func TestMinHeap(t *testing.T) {
	heap := &Heap[int]{}

	expected := 0
	got := heap.length
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	heap.insert(5)
	heap.insert(3)
	heap.insert(69)
	heap.insert(420)
	heap.insert(4)
	heap.insert(1)
	heap.insert(8)
	heap.insert(7)

	expected = 8
	got = heap.length
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 1
	got, _ = heap.pull()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 3
	got, _ = heap.pull()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 4
	got, _ = heap.pull()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 5
	got, _ = heap.pull()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 4
	got = heap.length
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 7
	got, _ = heap.pull()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 8
	got, _ = heap.pull()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 69
	got, _ = heap.pull()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 420
	got, _ = heap.pull()
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = 0
	got = heap.length
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

	_, err := heap.pull()
	if err == nil {
		t.Error("expected err, got nil")
	}

}
