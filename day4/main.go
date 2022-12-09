package main

import (
	"adventOfCode2022/pkg/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMinMax(rng string) (int64, int64, error) {
	split := strings.Split(rng, "-")
	min, err1 := strconv.ParseInt(split[0], 10, 64)
	max, err2 := strconv.ParseInt(split[1], 10, 64)
	var err error
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}

	return min, max, err
}

func hasFullyContainedRange(ranges [][]int64) bool {
	for i := 0; i < len(ranges); i++ {
		min := ranges[i][0]
		max := ranges[i][1]
		for j := i + 1; j < len(ranges); j++ {
			minj := ranges[j][0]
			maxj := ranges[j][1]

			if (min >= minj && max <= maxj) || (minj >= min && maxj <= max) {
				return true
			}
		}
	}

	return false
}

func processLine(in <-chan string) <-chan [][]int64 {
	output := make(chan [][]int64)

	go func() {
		defer close(output)
		for line := range in {
			rangesStr := strings.Split(line, ",")
			ranges := make([][]int64, len(rangesStr))
			for i, s := range rangesStr {
				min, max, err := getMinMax(s)
				if err != nil {
					continue
				}
				ranges[i] = []int64{min, max}
			}

			out := make([][]int64, len(ranges))
			copy(out, ranges)
			output <- out

		}
	}()

	return output
}

func hasOverlap(ranges [][]int64) bool {
	for i := 0; i < len(ranges); i++ {
		min := ranges[i][0]
		max := ranges[i][1]
		for j := i + 1; j < len(ranges); j++ {
			minj := ranges[j][0]
			maxj := ranges[j][1]

			if (min >= minj && min <= maxj) || (minj >= min && minj <= max) {
				return true
			}

		}
	}

	return false
}

func main() {
	args := os.Args[1:]

	var total int64
	for ranges := range processLine(util.ReadInputFileLines(args[0])) {
		if hasOverlap(ranges) {
			total++
		}
	}

	fmt.Printf("Total Fully Contained Ranges %d\n", total)
}
