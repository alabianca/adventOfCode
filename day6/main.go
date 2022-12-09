package main

import (
	"adventOfCode2022/pkg/ds/queue"
	"bufio"
	"fmt"
	"log"
	"os"
)

type FrequencyQueue struct {
	data      *queue.Queue[byte]
	frequency map[byte]int
}

func (f *FrequencyQueue) Enqueue(val byte) {
	f.data.EnqueueBack(val)
	if count, ok := f.frequency[val]; ok {
		f.frequency[val] = count + 1
	} else {
		f.frequency[val] = 1
	}
}

func (f *FrequencyQueue) Dequeue() (byte, error) {
	b, err := f.data.DeQueueFront()
	if err != nil {
		return b, err
	}

	if count, ok := f.frequency[b]; ok {
		if count == 1 {
			delete(f.frequency, b)
		} else {
			f.frequency[b]--
		}

	}

	return b, nil
}

func (f *FrequencyQueue) HasDuplicates() bool {
	return !(len(f.frequency) == f.data.Size())
}

func (f *FrequencyQueue) Size() int {
	return f.data.Size()
}

func readInput(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error %s\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	bytesRead := 0
	markerLength := 14
	fqueue := FrequencyQueue{
		frequency: make(map[byte]int),
		data:      queue.New[byte](),
	}

	for scanner.Scan() {
		byt := scanner.Bytes()[0]
		bytesRead++
		fmt.Printf("Enque %s\n", string(byt))
		fqueue.Enqueue(byt)

		if fqueue.Size() == markerLength {
			if !fqueue.HasDuplicates() {
				return bytesRead
			} else {
				b, _ := fqueue.Dequeue()
				fmt.Printf("Dequeued %s\n", string(b))
			}
		}
	}

	return bytesRead
}

func main() {
	args := os.Args[1:]
	fmt.Println(readInput(args[0]))
}
