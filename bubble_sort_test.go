package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

    unsorted := []int{9, 3, 7, 4, 69, 420, 42}
    sorted := []int{3, 4, 7, 9, 42, 69, 420}

    got := BubbleSort(unsorted)
    if !reflect.DeepEqual(sorted, got){
        t.Errorf("expected %v, got %v", sorted, got)
    }

}

