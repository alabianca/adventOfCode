package heap

import "fmt"

const EmptyHeapError = "empty heap"

// IntHeap Accessing
// values[(i-1)/2] => Parent Node
func parentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

// values[(i*2)+1] => Left Child Node
func leftChildIndex(parentIndex int) int {
	return (parentIndex * 2) + 1
}

// values[(i*2)+2] => Right Child Node
func rightChildIndex(parentIndex int) int {
	return leftChildIndex(parentIndex) + 1
}

type CompareFunc func(x, y int64) bool

func MaxHeap(x, y int64) bool {
	return x > y
}
func MinHeap(x, y int64) bool {
	return x < y
}

type IntHeap struct {
	values  []int64
	compare CompareFunc
	size    int
}

func NewIntHeap(capacity int, compare CompareFunc) *IntHeap {
	return &IntHeap{
		values:  make([]int64, 0, capacity),
		compare: compare,
		size:    0,
	}
}

func (heap *IntHeap) swap(parentIndex, childIndex int) {
	temp := heap.values[childIndex]
	heap.values[childIndex] = heap.values[parentIndex]
	heap.values[parentIndex] = temp
}

func (heap *IntHeap) pop() (int64, error) {
	if len(heap.values) == 0 {
		return 0, fmt.Errorf(EmptyHeapError)
	}
	item := heap.values[len(heap.values)-1]
	heap.values = append(heap.values[:heap.size-1])
	heap.size--

	return item, nil
}

func (heap *IntHeap) hasLeftChild(index int) bool {
	return leftChildIndex(index) < len(heap.values)
}
func (heap *IntHeap) hasRightChild(index int) bool {
	return rightChildIndex(index) < len(heap.values)
}

func (heap *IntHeap) Peek() int64 {
	return heap.values[0]
}

func (heap *IntHeap) RemoveRoot() (int64, error) {
	if len(heap.values) == 0 {
		return 0, fmt.Errorf(EmptyHeapError)
	}
	item := heap.Peek()
	heap.values[0] = heap.values[len(heap.values)-1]
	heap.pop()

	// fix heap violations down starting at the root
	index := 0
	for heap.hasLeftChild(index) {
		// check which of the children is smaller/larger (depending on compare function)
		childIndex := leftChildIndex(index)
		if heap.hasRightChild(index) && heap.compare(heap.values[rightChildIndex(index)], heap.values[childIndex]) {
			childIndex = rightChildIndex(index)
		}

		// check if child is smaller/larger than parent. If not stop
		if heap.compare(heap.values[index], heap.values[childIndex]) {
			break
		}

		heap.swap(index, childIndex)
		index = childIndex
	}

	return item, nil
}

func (heap *IntHeap) Insert(v int64) bool {
	// add element ad the end of heap
	index := heap.size
	heap.values = append(heap.values, v)
	heap.size++

	// fix heap violations
	for index > 0 && !heap.compare(heap.values[parentIndex(index)], heap.values[index]) {
		heap.swap(parentIndex(index), index)
		index = parentIndex(index)

	}

	return true
}
