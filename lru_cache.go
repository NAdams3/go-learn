package main

import (
	"errors"
	"hash"
	"hash/fnv"
)

type LRUCache[V comparable] struct {
	data   [][]*element[KeyValue[V]]
	queue  Queue[string]
	hasher hash.Hash32
}

type KeyValue[V comparable] struct {
	key   string
	value V
}

func (c *LRUCache[V]) get(key string) (V, error) {

	if c.queue.length == 0 {
		return *new(V), errors.New("No items in cache")
	}

	index := c.hash(key)
	if len(c.data[index]) == 0 {
		return *new(V), errors.New("Key Value not found")
	}

	found := false
	val := *new(V)
	for _, elem := range c.data[index] {
		if elem.value.key == key {
			found = true
			val = elem.value.value
			break
		}
	}

	if found {
		return val, nil
	}

	return *new(V), errors.New("Key value not found")

}

func (c *LRUCache[V]) update(key string, value V) {

	index := c.hash(key)
	keyValue := KeyValue[V]{
		key:   key,
		value: value,
	}

	if len(c.data) == 0 {
		c.data = make([][]KeyValue[V], 3)
	}

	c.data[index] = append(c.data[index], keyValue)
	c.queue.enqueue(key)

}

func (c *LRUCache[V]) hash(key string) int {
	if c.hasher == nil {
		c.hasher = fnv.New32a()
	}
	c.hasher.Write([]byte(key))
	return int(c.hasher.Sum32()) % 3

}
