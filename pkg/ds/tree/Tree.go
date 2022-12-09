package tree

type Tree[K comparable, T comparable] struct {
	Parent   *Tree[K, T]
	Children map[K]*Tree[K, T]
	Value    T
}

func New[K comparable, T comparable](parent *Tree[K, T], val T) *Tree[K, T] {
	return &Tree[K, T]{
		Parent:   parent,
		Children: make(map[K]*Tree[K, T]),
		Value:    val,
	}
}

func (t *Tree[K, T]) GetChild(k K) (*Tree[K, T], bool) {
	child, ok := t.Children[k]
	return child, ok

}

func (t *Tree[K, T]) Insert(k K, v T) {
	t.Children[k] = New[K, T](t, v)
}
