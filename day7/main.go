package main

import (
	"adventOfCode2022/pkg/ds/tree"
	"adventOfCode2022/pkg/util"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const commandPrefix = "$"
const dirPrefix = "dir"
const cdCommand = "cd"
const lsCommand = "ls"
const cdBack = ".."

func currentDirName(dir *tree.Tree[string, int]) string {
	parent := dir.Parent
	if parent == nil {
		return "root"
	}
	for k, v := range parent.Children {
		if v == dir {
			return k
		}
	}

	return "root"
}

func traverseDF(t *tree.Tree[string, int], fn func(k string, v int)) {
	for k, child := range t.Children {
		fn(k, child.Value)
		traverseDF(child, fn)
	}
}

func traverseDirectories(name string, root *tree.Tree[string, int], min, max int64, projection func(name string, size int)) int {
	currentSize := root.Value

	for k, child := range root.Children {
		currentSize += traverseDirectories(k, child, min, max, projection)
	}

	// this is a directory check if it meets the projection
	if root.Value == 0 && int64(currentSize) >= min && int64(currentSize) <= max {
		projection(name, currentSize)
	}

	return currentSize

}

func getSize(root *tree.Tree[string, int]) int {
	currentSize := root.Value

	for _, child := range root.Children {
		currentSize += getSize(child)
	}

	return currentSize
}

func processInput(in <-chan string) (int, error) {
	root := tree.New[string, int](nil, 0)
	// insert the root directory
	root.Insert("/", 0)

	current := root
	var currentCommand string

	for line := range in {
		split := strings.Split(line, " ")
		switch split[0] {
		case commandPrefix:
			currentCommand = split[1]
			if currentCommand == cdCommand {
				//fmt.Printf("changing directory from %s to %s\n", currentDirName(current), split[2])
				if child, ok := current.GetChild(split[2]); ok {
					current = child
				} else if split[2] == cdBack {
					current = current.Parent
				} else {
					return 0, fmt.Errorf("directory %s not found in %s", split[2], currentDirName(current))

				}
			}

		default:
			var size int64
			if split[0] != dirPrefix {
				size, _ = strconv.ParseInt(split[0], 10, 64)
			}
			current.Insert(split[1], int(size))
		}
	}

	rootDir, _ := root.GetChild("/")
	maxSpaceAvailable := 70000000
	usedSpace := getSize(rootDir)
	availableSpace := maxSpaceAvailable - usedSpace
	smallest := usedSpace
	updateSize := 30000000
	sizeNeeded := updateSize - availableSpace

	traverseDirectories("/", rootDir, int64(sizeNeeded), math.MaxInt64, func(name string, size int) {
		fmt.Printf("found dir %s => %d\n", name, size)
		smallest = int(math.Min(float64(smallest), float64(size)))
	})

	return smallest, nil
}

func main() {
	args := os.Args[1:]
	total, _ := processInput(util.ReadInputFileLines(args[0]))
	fmt.Printf("total %d\n", total)
}
