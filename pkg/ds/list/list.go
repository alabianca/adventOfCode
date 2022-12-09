package list

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	head *Node[T]
	//tail *Node[T]
	size int
}

func New[T any]() *List[T] {
	return &List[T]{
		head: nil,
	}
}

// Push an item to the end of the list
func (l *List[T]) Push(val T) {
	l.size++
	if l.head == nil {
		l.head = newNode(nil, val)
		return
	}

	node := l.HeadNode()
	for node.Next() != nil {
		node = node.Next()
	}
	node.next = newNode[T](node, val)
}

// Unshift adds an item to the front of the list
func (l *List[T]) Unshift(val T) {
	l.size++
	if l.head == nil {
		l.head = newNode(nil, val)
		return
		//l.tail = l.head
	} else {
		head := l.head
		l.head = newNode(nil, val)
		head.prev = l.head
		l.head.next = head
	}
}

// Shift removes the head and returns its value
func (l *List[T]) Shift() (T, error) {
	node := l.HeadNode()
	if node == nil {
		var noop T
		return noop, fmt.Errorf("empty list")
	}
	l.size--
	l.head = node.next
	node.next = nil

	return node.Value(), nil
}

// Pop removes the last element of the list and returns its value
func (l *List[T]) Pop() (T, error) {
	node := l.HeadNode()
	if node == nil {
		var noop T
		return noop, fmt.Errorf("empty list")
	}

	slow := node
	fast := node.Next()

	for fast != nil && fast.Next() != nil {
		slow = slow.Next()
		fast = fast.Next()
	}

	slow.next = nil

	node.prev = nil

	return node.Value(), nil
}

func (l *List[T]) Splice(startIndex int, endIndex int) *List[T] {
	if l.Size() == 0 {
		return &List[T]{}
	}
	curr := l.head
	head := curr
	tail := curr
	end := endIndex
	if end > l.size {
		end = l.size
	}

	for i := 1; i < end; i++ {
		curr = curr.Next()
		if i == startIndex {
			head = curr
		}
	}
	tail = curr

	if head.prev == nil {
		l.head = nil
		l.size = 0
	} else {
		head.prev.next = tail.next
		l.size = l.size - (end - startIndex)
	}

	head.prev = nil
	tail.next = nil

	return &List[T]{
		head: head,
	}

}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Head() (T, error) {
	if l.head == nil {
		var noop T
		return noop, fmt.Errorf("empty list")
	}
	return l.head.val, nil
}

func (l *List[T]) HeadNode() *Node[T] {
	return l.head
}

func (l List[T]) String() string {
	builder := strings.Builder{}
	builder.WriteString("[")
	for n := l.HeadNode(); n != nil; n = n.Next() {
		builder.WriteString(fmt.Sprintf("%v->", n.Value()))
	}
	builder.WriteString("]")

	return builder.String()
}

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]
	val  T
}

func newNode[T any](prev *Node[T], v T) *Node[T] {
	return &Node[T]{
		next: nil,
		prev: prev,
		val:  v,
	}
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

func (n *Node[T]) Value() T {
	return n.val
}
