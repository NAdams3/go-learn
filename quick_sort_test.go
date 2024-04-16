package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
    unsorted := []int{9, 3, 7, 4, 69, 420, 42}
    sorted := []int{3, 4, 7, 9, 42, 69, 420}

    got := QuickSort(unsorted, len(unsorted) - 1, 0)
    if !reflect.DeepEqual(unsorted, sorted) {
        t.Errorf("expected %v, got %v", sorted, got)
    }

}

