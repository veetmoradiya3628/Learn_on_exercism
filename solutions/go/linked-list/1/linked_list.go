package linkedlist

import "errors"

type Node struct {
	Value any
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func NewList(elements ...any) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

func (l *List) Unshift(v any) {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head.prev = n
		l.head = n
	}
	l.size++
}

func (l *List) Push(v any) {
	n := &Node{Value: v}
	if l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
	l.size++
}

func (l *List) Shift() (any, error) {
	if l.head == nil {
		return nil, errors.New("cannot shift from an empty list")
	}

	val := l.head.Value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	l.size--
	return val, nil
}

func (l *List) Pop() (any, error) {
	if l.tail == nil {
		return nil, errors.New("cannot pop from an empty list")
	}

	val := l.tail.Value
	l.tail = l.tail.prev

	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	l.size--
	return val, nil
}

func (l *List) Reverse() {
	curr := l.head
	l.head, l.tail = l.tail, l.head

	for curr != nil {
		curr.next, curr.prev = curr.prev, curr.next
		curr = curr.prev
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Count() int {
	return l.size
}

func (l *List) Delete(v any) bool {
	curr := l.head
	for curr != nil {
		if curr.Value == v {
			if curr.prev != nil {
				curr.prev.next = curr.next
			} else {
				l.head = curr.next
			}

			if curr.next != nil {
				curr.next.prev = curr.prev
			} else {
				l.tail = curr.prev
			}

			l.size--
			return true
		}
		curr = curr.next
	}
	return false
}
