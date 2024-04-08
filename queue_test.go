package main

import "testing"

func TestQueue(t *testing.T) {
    list := Queue[int]{}

    list.enqueue(5)
    list.enqueue(7)
    list.enqueue(9)

    got, err := list.deque()
    if err != nil {
        t.Error("expected 5, got error")
    }

    if got != 5 {
        t.Errorf("expected %d, got %d", 5, got)
    }

    if list.length != 2 {
        t.Errorf("expected 2, got %d", list.length)
    }

    list.enqueue(11)

    got, _ = list.deque()
    if got != 7 {
        t.Errorf("expected %d, got %d", 7, got)
    }

    got, _ = list.deque()
    if got != 9 {
        t.Errorf("expected %d, got %d", 9, got)
    }

    got, err = list.peek()
    if err != nil {
        t.Error("expected 11, got error")
    }

    if got != 11 {
        t.Errorf("expected %d, got %d", 11, got)
    }

    got, _ = list.deque()
    if got != 11 {
        t.Errorf("expected %d, got %d", 11, got)
    }

    _, err = list.deque()
    if err == nil {
        t.Error("expected error, got nil")
    }

    if list.length != 0 {
        t.Errorf("expected 0, got %d", list.length)
    }

    list.enqueue(69)

    got, _ = list.peek()
    if got != 69 {
        t.Errorf("expected %d, got %d", 69, got)
    }
    
    if list.length != 1 {
        t.Errorf("expected 1, got %d", list.length)
    }

}

