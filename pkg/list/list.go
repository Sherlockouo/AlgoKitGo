package list

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T any]() *List[T] {
	return &List[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *List[T]) Len() int {
	return l.size
}

func (l *List[T]) Front() *Node[T] {
	return l.head
}

func (l *List[T]) Back() *Node[T] {
	return l.tail
}

func (l *List[T]) PushFront(value T) *Node[T] {
	node := &Node[T]{
		Value: value,
		Next:  l.head,
	}
	if l.head != nil {
		l.head.Prev = node
	} else {
		l.tail = node
	}
	l.head = node
	l.size++
	return node
}

func (l *List[T]) PushBack(value T) *Node[T] {
	node := &Node[T]{
		Value: value,
		Prev:  l.tail,
	}
	if l.tail != nil {
		l.tail.Next = node
	} else {
		l.head = node
	}
	l.tail = node
	l.size++
	return node
}

func (l *List[T]) InsertAfter(node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	}
	newNode := &Node[T]{
		Value: value,
		Prev:  node,
		Next:  node.Next,
	}
	if node.Next != nil {
		node.Next.Prev = newNode
	} else {
		l.tail = newNode
	}
	node.Next = newNode
	l.size++
	return newNode
}

func (l *List[T]) Index(index int) *Node[T] {
	if index < 0 || index >= l.size {
		return nil
	}
	if index < l.size/2 {
		node := l.head
		for range index {
			node = node.Next
		}
		return node
	} else {
		node := l.tail
		for i := l.size - 1; i > index; i-- {
			node = node.Prev
		}
		return node
	}
}

func (l *List[T]) Remove(node *Node[T]) {
	if node == nil {
		return
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.tail = node.Prev
	}
	l.size--
}
