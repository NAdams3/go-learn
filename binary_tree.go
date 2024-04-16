package main

type BinaryNode[T comparable] struct {
	value T
	right *BinaryNode[T]
	left  *BinaryNode[T]
}

func (head *BinaryNode[T]) PreOrderSearch() []T {
	values := make([]T, 0)
	head.preOrderWalk(&values)

	return values
}

func (node *BinaryNode[T]) preOrderWalk(path *[]T) {
	if node == nil {
		return
	}

	*path = append(*path, node.value)

	node.left.preOrderWalk(path)
	node.right.preOrderWalk(path)
}

func (head *BinaryNode[T]) InOrderSearch() []T {
	values := make([]T, 0)
	head.inOrderWalk(&values)

	return values
}

func (node *BinaryNode[T]) inOrderWalk(path *[]T) {
	if node == nil {
		return
	}

	node.left.inOrderWalk(path)

	*path = append(*path, node.value)

	node.right.inOrderWalk(path)
}
