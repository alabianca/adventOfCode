package heap

import (
	"testing"
)

func TestIntHeap_Insert_MaxHeap(t *testing.T) {
	heap := NewIntHeap(5, MaxHeap)

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(17)
	heap.Insert(10)
	heap.Insert(84)
	heap.Insert(19)
	heap.Insert(6)
	heap.Insert(22)
	heap.Insert(9)

	order := [9]int{84, 22, 19, 17, 10, 9, 6, 5, 3}

	for _, item := range order {
		x, err := heap.RemoveRoot()
		if err != nil {
			t.Errorf("Expected no error, but got %s\n", err)
		}
		if x != int64(item) {
			t.Errorf("Expected %d, but got %d\n", item, x)
		}
	}
}

func TestIntHeap_Insert_MinHeap(t *testing.T) {
	heap := NewIntHeap(5, MinHeap)

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(17)
	heap.Insert(10)
	heap.Insert(84)
	heap.Insert(19)
	heap.Insert(6)
	heap.Insert(22)
	heap.Insert(9)

	order := [9]int{3, 5, 6, 9, 10, 17, 19, 22, 84}

	for _, item := range order {
		x, err := heap.RemoveRoot()
		if err != nil {
			t.Errorf("Expected no error, but got %s\n", err)
		}
		if x != int64(item) {
			t.Errorf("Expected %d, but got %d\n", item, x)
		}
	}
}
