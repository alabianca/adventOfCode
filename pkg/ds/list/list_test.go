package list

import "testing"

func TestList_Splice(t *testing.T) {
	myList := New[int]()

	myList.Push(1)
	myList.Push(2)
	myList.Push(3)
	myList.Push(4)
	myList.Push(5)
	myList.Push(6)

	splice := myList.Splice(1, 3)
	expectedSplice := []int{2, 3}
	j := 0
	for n := splice.HeadNode(); n != nil; n = n.Next() {
		if n.Value() != expectedSplice[j] {
			t.Errorf("Expected %d, but got %d\n", expectedSplice[j], n.Value())
		}
		j++
	}

	expected := []int{1, 4, 5, 6}
	i := 0
	for n := myList.HeadNode(); n != nil; n = n.Next() {
		if n.Value() != expected[i] {
			t.Errorf("Expected %d, but got %d\n", expected[i], n.Value())
		}
		i++
	}

	if myList.Size() != 4 {
		t.Errorf("Expected size %d after splice, but got %d\n", 4, myList.Size())
	}
}

func TestList_Splice_SingeItem(t *testing.T) {
	myList := New[int]()

	myList.Push(1)

	splice := myList.Splice(0, 3)
	expectedSplice := []int{1}
	j := 0
	for n := splice.HeadNode(); n != nil; n = n.Next() {
		if n.Value() != expectedSplice[j] {
			t.Errorf("Expected %d, but got %d\n", expectedSplice[j], n.Value())
		}
		j++
	}

	if myList.HeadNode() != nil {
		t.Errorf("Expected nil, but got %d", myList.HeadNode().Value())
	}

	if myList.Size() != 0 {
		t.Errorf("Expected size %d after splice, but got %d\n", 0, myList.Size())
	}
}
