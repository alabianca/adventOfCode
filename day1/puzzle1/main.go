package main

import (
	"adventOfCode2022/pkg/ds/heap"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Printf("input file is required")
		os.Exit(1)
	}

	elfs := readInput(args[0])

	maxHeap := heap.NewIntHeap(10, heap.MaxHeap)
	for e := range elfs {
		maxHeap.Insert(e.calories)
	}

	fmt.Printf("Elf with with calories %d\n", maxHeap.Peek())
	var total int64
	for i := 0; i < 3; i++ {
		x, _ := maxHeap.RemoveRoot()
		total += x
	}
	fmt.Printf("The top 3 elfs have a total of %d calories\n", total)
}
