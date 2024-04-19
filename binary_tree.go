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

func (head *BinaryNode[T]) BreadthFirstSearch(needle T) (bool, error) {

	found := false

	queue := Queue[*BinaryNode[T]]{}
	queue.enqueue(head)

	for queue.length > 0 {
		node, err := queue.deque()
		if err != nil {
			return false, err
		}

		if node.value == needle {
			found = true
			break
		}

		if node.left != nil {
			queue.enqueue(node.left)
		}

		if node.right != nil {
			queue.enqueue(node.right)
		}

	}

	return found, nil
}

func Compare[T comparable](head1, head2 *BinaryNode[T]) bool {
	isEqual := true
	compareWalk(head1, head2, &isEqual)
	return isEqual
}

func compareWalk[T comparable](node1, node2 *BinaryNode[T], isEqual *bool) {
	if !*isEqual {
		return
	}

	if node1 == nil && node2 == nil {
		return
	}

	if node1 != nil && node2 != nil {
		if node1.value != node2.value {
			*isEqual = false
			return
		}
	} else {
		*isEqual = false
		return
	}

	compareWalk(node1.left, node2.left, isEqual)
	compareWalk(node1.right, node2.right, isEqual)
}
