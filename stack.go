package main

import (
	"errors"
)

type Stack[T comparable] struct {
	head, tail *element[T]
	length     int
}

func (s *Stack[T]) push(val T) {
	item := &element[T]{
		value: val,
		next:  s.head,
	}

	if s.length == 0 {
		s.head = item
		s.tail = item
	} else {
		s.head = item
	}

	s.length++
}

func (s *Stack[T]) pop() (T, error) {
	if s.length == 0 {
		return *new(T), errors.New("No items in stack")
	}

	val := s.head.value
	s.head = s.head.next
	s.length--

	if s.head == nil {
		s.tail = nil
	}

	return val, nil

}

func (s *Stack[T]) peek() (T, error) {
	if s.length == 0 {
		return *new(T), errors.New("No items in stack")
	}
	return s.head.value, nil
}
