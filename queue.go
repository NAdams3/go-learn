package main

import "errors"

type Queue[T any] struct {
    head, tail *element[T]
    length int
}

type element[T any] struct {
    value T
    next *element[T]
}

func (q *Queue[T]) enqueue(val T) {
    item := &element[T]{
        value: val,
        next: nil,
    }

    if q.length == 0 {
        q.head = item
        q.tail = item
    } else {
        q.tail.next = item
        q.tail = item
    }

    q.length++
}

func (q *Queue[T]) deque() (T, error) {
    if q.length == 0 {
        return *new(T), errors.New("No items in queue")
    }

    value := q.head.value
    q.head = q.head.next
    q.length--

    if q.head == nil {
        q.tail = nil 
    }

    return value, nil
    
}

func (q *Queue[T]) peek() (T, error) {
    if q.length == 0 {
        return *new(T), errors.New("No items in queue")
    }
    return q.head.value, nil
}

