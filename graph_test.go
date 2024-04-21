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

type WeightedEdge struct {
	to, weight int
}

var list = [][]WeightedEdge{
	{WeightedEdge{to: 1, weight: 3}, WeightedEdge{to: 2, weight: 1}},
	{WeightedEdge{to: 0, weight: 3}, WeightedEdge{to: 2, weight: 4}, WeightedEdge{to: 4, weight: 1}},
	{WeightedEdge{to: 1, weight: 4}, WeightedEdge{to: 3, weight: 7}, WeightedEdge{to: 0, weight: 1}},
	{WeightedEdge{to: 6, weight: 1}, WeightedEdge{to: 2, weight: 7}, WeightedEdge{to: 4, weight: 5}},
	{WeightedEdge{to: 5, weight: 2}, WeightedEdge{to: 3, weight: 5}, WeightedEdge{to: 1, weight: 1}},
	{WeightedEdge{to: 6, weight: 1}, WeightedEdge{to: 2, weight: 18}, WeightedEdge{to: 4, weight: 2}},
	{WeightedEdge{to: 3, weight: 1}, WeightedEdge{to: 5, weight: 1}},
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

func TestGraphDijkstraList(t *testing.T) {

	expected := []int{0, 1, 4, 5, 6}
	got := DijkstraList(list, 0, 6)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
