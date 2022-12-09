package stack

import "adventOfCode2022/pkg/ds/list"

type Stack[T any] struct {
	data *list.List[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		data: list.New[T](),
	}
}

func (stack *Stack[T]) Push(v T) {
	stack.data.Push(v)
}

func (stack *Stack[T]) Pop() (T, error) {
	return stack.data.Pop()
}
