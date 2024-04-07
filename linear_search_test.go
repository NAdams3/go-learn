package main

import (
    "testing"
)

func TestLinearSearch(t *testing.T) {
    test := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420};

    got := LinearSearch(test, 69)
    if got != true {
        t.Errorf("Expected true, got %t", got)
    }

    got = LinearSearch(test, 1336)
    if got != false {
        t.Errorf("Expected false, got %t", got)
    }

    got = LinearSearch(test, 69420)
    if got != true {
        t.Errorf("Expected true, got %t", got)
    }

    got = LinearSearch(test, 69421)
    if got != false {
        t.Errorf("Expected false, got %t", got)
    }
    
    got = LinearSearch(test, 1)
    if got != true {
        t.Errorf("Expected true, got %t", got)
    }

    got = LinearSearch(test, 0)
    if got != false {
        t.Errorf("Expected false, got %t", got)
    }

}

