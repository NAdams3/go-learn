package main

import (
	"reflect"
	"testing"
)

var testBinaryNode = &BinaryNode[int]{
	value: 20,
	right: &BinaryNode[int]{
		value: 50,
		right: &BinaryNode[int]{
			value: 100,
			right: nil,
			left:  nil,
		},
		left: &BinaryNode[int]{
			value: 30,
			right: &BinaryNode[int]{
				value: 45,
				right: nil,
				left:  nil,
			},
			left: &BinaryNode[int]{
				value: 29,
				right: nil,
				left:  nil,
			},
		},
	},
	left: &BinaryNode[int]{
		value: 10,
		right: &BinaryNode[int]{
			value: 15,
			right: nil,
			left:  nil,
		},
		left: &BinaryNode[int]{
			value: 5,
			right: &BinaryNode[int]{
				value: 7,
				right: nil,
				left:  nil,
			},
			left: nil,
		},
	},
}

func TestBinaryTreePreOrderSearch(t *testing.T) {
	expected := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	got := testBinaryNode.PreOrderSearch()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestBinaryTreeInOrderSearch(t *testing.T) {
	expected := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	got := testBinaryNode.InOrderSearch()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
