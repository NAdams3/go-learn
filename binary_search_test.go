package main

import (
    "testing"
)

func TestBinarySearch(t *testing.T) {
    test := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420};

    got := BinarySearch(test, 69)
    if got != true {
        t.Errorf("Expected true, got %t", got)
    }

    got = BinarySearch(test, 1336)
    if got != false {
        t.Errorf("Expected false, got %t", got)
    }

    got = BinarySearch(test, 69420)
    if got != true {
        t.Errorf("Expected true, got %t", got)
    }

    got = BinarySearch(test, 69421)
    if got != false {
        t.Errorf("Expected false, got %t", got)
    }
    
    got = BinarySearch(test, 1)
    if got != true {
        t.Errorf("Expected true, got %t", got)
    }

    got = BinarySearch(test, 0)
    if got != false {
        t.Errorf("Expected false, got %t", got)
    }

}

func TestBinarySearchRecursive(t *testing.T) {
    test := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420};

    got := BinarySearchRecursive(test, 69)
    if got != true {
        t.Errorf("Expected true, got %t", got)
    }

    got = BinarySearchRecursive(test, 1336)
    if got != false {
        t.Errorf("Expected false, got %t", got)
    }

    got = BinarySearchRecursive(test, 69420)
    if got != true {
        t.Errorf("Expected true, got %t", got)
    }

    got = BinarySearchRecursive(test, 69421)
    if got != false {
        t.Errorf("Expected false, got %t", got)
    }
    
    got = BinarySearchRecursive(test, 1)
    if got != true {
        t.Errorf("Expected true, got %t", got)
    }

    got = BinarySearchRecursive(test, 0)
    if got != false {
        t.Errorf("Expected false, got %t", got)
    }

}


