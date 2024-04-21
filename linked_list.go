package main

import (
	"errors"
	"math"
)

type LinkedList[T comparable] struct {
	head, tail *element[T]
	length     int
}

func (list *LinkedList[T]) insertAt(val T, index int) error {

	currentNode, err := list.getElement(index)
	if err != nil {
		return err
	}

	node := &element[T]{
		value: val,
		next:  currentNode,
		prev:  currentNode.prev,
	}

	currentNode.prev.next = node
	currentNode.prev = node

	list.length++

	return nil

}

func (list *LinkedList[T]) remove(val T) (T, error) {
	found := false
	index := 0

	current := list.head
	for i := 0; i < list.length; i++ {
		if current.value == val {
			found = true
			index = i
			break
		}
		current = current.next
	}

	if found == false {
		return *new(T), errors.New("Node with value not found")
	}

	if index == 0 {
		list.head = current.next
	}
	if index == list.length-1 {
		list.tail = current.prev
	}

	if current.prev != nil {

		current.prev.next = current.next
	}
	if current.next != nil {
		current.next.prev = current.prev
	}

	list.length--

	return current.value, nil

}

func (list *LinkedList[T]) removeAt(index int) (T, error) {

	current, err := list.getElement(index)
	if err != nil {
		return *new(T), err
	}

	if index == 0 {
		list.head = current.next
	}

	if index == list.length-1 {
		list.tail = current.prev
	}

	if current.prev != nil {

		current.prev.next = current.next
	}
	if current.next != nil {
		current.next.prev = current.prev
	}
	list.length--

	return current.value, nil

}

func (list *LinkedList[T]) append(val T) {
	node := &element[T]{
		value: val,
		next:  nil,
		prev:  list.tail,
	}

	if list.length == 0 {
		list.head = node
	} else {
		list.tail.next = node
	}

	list.tail = node
	list.length++
}

func (list *LinkedList[T]) prepend(val T) {
	node := &element[T]{
		value: val,
		next:  list.head,
		prev:  nil,
	}

	if list.length > 0 {
		list.head.prev = node
	} else {
		list.tail = node
	}

	list.head = node
	list.length++
}

func (list *LinkedList[T]) get(index int) (T, error) {
	current, err := list.getElement(index)
	if err != nil {
		return *new(T), err
	}

	return current.value, nil
}

func (list *LinkedList[T]) getElement(index int) (*element[T], error) {
	if index < 0 || index >= list.length {
		return nil, errors.New("Index out of range")
	}

	if index == 0 {
		return list.head, nil
	}

	if index == list.length-1 {
		return list.tail, nil
	}

	midIndex := int(math.Floor(float64((list.length - 1) / 2)))
	current := list.head

	if index > midIndex {
		current = list.tail
		for i := list.length - 1; i > index; i-- {
			current = current.prev
		}
		return current, nil
	}

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current, nil

}
