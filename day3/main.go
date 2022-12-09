package main

import (
	"adventOfCode2022/pkg/util"
	"fmt"
	"os"
)

func toPriority(b byte) int {
	// upper case
	if b < byte('a') {
		return int(b-byte('A')) + 27
	}
	// lower case
	return int(b-byte('a')) + 1
}

func findMatchWithPriority(c1, c2 []byte) int {
	map1 := make(map[byte]bool)
	map2 := make(map[byte]bool)

	for _, b := range c1 {
		map1[b] = true
	}
	for _, b := range c2 {
		map2[b] = true
	}

	var found byte
	for b, _ := range map1 {
		if _, ok := map2[b]; ok {
			found = b
			break
		}
	}

	return toPriority(found)

}

func groupByX(in <-chan string, groupSize int) <-chan []string {
	output := make(chan []string)

	go func() {
		defer close(output)
		group := make([]string, groupSize)
		index := 0
		for s := range in {
			group[index] = s
			index++
			if index == groupSize {
				out := make([]string, groupSize)
				copy(out, group)
				output <- out
				index = 0
			}

		}
	}()

	return output
}

func findCommonPriority(frequencyMap map[byte]int, numberOfElfs int) int {
	var found byte
	for b, v := range frequencyMap {
		if v == numberOfElfs {
			fmt.Printf("Found %s\n", string(b))
			found = b
			break
		}
	}
	return toPriority(found)
}

func insert(frequencyMap map[byte]int, b []byte) {
	for _, bt := range b {
		if count, ok := frequencyMap[bt]; ok {
			frequencyMap[bt] = count + 1
		} else {
			frequencyMap[bt] = 1
		}
	}

}

func createFrequencyMap(b []byte) []int {
	fmap := make([]int, 52)

	for _, bt := range b {
		priority := toPriority(bt) - 1
		if fmap[priority] > 0 {
			fmap[priority]++
		} else {
			fmap[priority] = 1
		}
	}

	return fmap
}

func main() {
	args := os.Args[1:]
	var total int
	for rs := range groupByX(util.ReadInputFileLines(args[0]), 3) {
		fmap1 := createFrequencyMap([]byte(rs[0]))
		fmap2 := createFrequencyMap([]byte(rs[1]))
		fmap3 := createFrequencyMap([]byte(rs[2]))

		for i := 0; i < 52; i++ {
			if fmap1[i] >= 1 && fmap2[i] >= 1 && fmap3[i] >= 1 {
				total = total + i + 1
			}
		}
	}

	fmt.Printf("Total %d\n", total)
}
