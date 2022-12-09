package queue

import (
	"adventOfCode2022/pkg/ds/list"
)

type Queue[T any] struct {
	data *list.List[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		data: list.New[T](),
	}
}

func (q *Queue[T]) EnqueueBack(val T) {
	q.data.Push(val)
}

func (q *Queue[T]) EnqueueBackList(val *list.List[T]) {
	for n := val.HeadNode(); n != nil; n = n.Next() {
		q.data.Push(n.Value())
	}

}

func (q *Queue[T]) EnqueueFront(val T) {
	q.data.Unshift(val)
}

func (q *Queue[T]) DeQueueBack() (T, error) {
	return q.data.Pop()
}

func (q *Queue[T]) DeQueueFront() (T, error) {
	return q.data.Shift()
}

func (q *Queue[T]) Size() int {
	return q.data.Size()
}

func (q *Queue[T]) DequeueRange(startIndex, endIndex int) *list.List[T] {
	return q.data.Splice(startIndex, endIndex)
}

func (q *Queue[T]) PeekFront() (T, bool) {
	head, err := q.data.Head()
	if err != nil {
		return head, false
	}
	return head, true
}

func (q *Queue[T]) PeekBack() (T, bool) {
	node := q.data.HeadNode()
	if node == nil {
		var noop T
		return noop, false
	}

	for node.Next() != nil {
		node = node.Next()
	}

	return node.Value(), true
}

func (q Queue[T]) String() string {
	return q.data.String()
}
