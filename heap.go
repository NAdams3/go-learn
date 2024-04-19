package main

import (
	"cmp"
	"errors"
)

// natodo add type and implement max heap
type Heap[T cmp.Ordered] struct {
	length int
	data   []T
}

func (h *Heap[T]) insert(val T) {
	h.data = append(h.data, val)
	h.heapifyUp(h.length)
	h.length++
}

func (h *Heap[T]) pull() (T, error) {

	if h.length == 0 {
		return *new(T), errors.New("Heap is empty")
	}

	val := h.data[0]
	h.data[0] = h.data[h.length-1]

	h.length--

	h.heapifyDown(0)

	return val, nil

}

func (h *Heap[T]) heapifyUp(index int) {

	if index == 0 {
		return
	}

	parentIndex := parent(index)
	parent := h.data[parentIndex]

	if h.data[index] >= parent {
		return
	}

	h.data[parentIndex] = h.data[index]
	h.data[index] = parent

	h.heapifyUp(parentIndex)

}

func (h *Heap[T]) heapifyDown(index int) {

	compareChildIndex, rChildIndex := children(index)

	if compareChildIndex >= h.length {
		return
	}
	compareChild := h.data[compareChildIndex]

	if rChildIndex < h.length {
		rChild := h.data[rChildIndex]

		if compareChild > rChild {
			compareChild = rChild
			compareChildIndex = rChildIndex
		}
	}

	if h.data[index] <= compareChild {
		return
	}

	h.data[compareChildIndex] = h.data[index]
	h.data[index] = compareChild

	h.heapifyDown(compareChildIndex)

}

func parent(index int) int {
	return (index - 1) / 2
}

func children(index int) (int, int) {
	return 2*index + 1, 2*index + 2
}
