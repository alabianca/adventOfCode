package main

import (
	"bufio"
	"os"
	"strconv"
)

type elf struct {
	calories int64
	index    int64
}

func readInput(path string) <-chan elf {
	output := make(chan elf)

	go func() {
		defer close(output)
		file, err := os.Open(path)
		if err != nil {
			return
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var index int64
		var current int64
		for scanner.Scan() {
			text := scanner.Text()
			if len(text) > 0 {
				x, err := strconv.ParseInt(text, 10, 64)
				if err != nil {
					continue
				}
				current += x
			} else {
				output <- elf{calories: current, index: index}
				current = 0
				index++
			}
		}
	}()

	return output
}
