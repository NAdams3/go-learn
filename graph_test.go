package main

import (
	"reflect"
	"testing"
)

var matrix = [][]int{
	{0, 3, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 7, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 5, 0, 2, 0},
	{0, 0, 18, 0, 0, 0, 1},
	{0, 0, 0, 1, 0, 0, 1},
}

func TestBreadthFirstSearchGraphMatrix(t *testing.T) {

	expected := []int{0, 1, 4, 5, 6}
	got := GraphMatrixBreadthFirstSearch(matrix, 0, 6)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}

	expected = []int{}
	got = GraphMatrixBreadthFirstSearch(matrix, 6, 0)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
