package main

import (
	"adventOfCode2022/pkg/util"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func processInput(path string) [][]int {
	grid := make([][]int, 0)
	index := 0
	for line := range util.ReadInputFileLines(path) {
		grid = append(grid, make([]int, len(line)))

		sheights := strings.Split(line, "")
		for i, h := range sheights {
			intH, _ := strconv.ParseInt(h, 10, 64)
			grid[index][i] = int(intH)
		}
		index++
	}

	return grid
}

func getTreeHeightAt(rowIndex, colIndex int, grid [][]int) int {
	// account for index out of bounds here. If out of bounds we return -1 (those are the edges)
	if rowIndex < 0 || rowIndex >= len(grid) || colIndex < 0 || colIndex >= len(grid[rowIndex]) {
		return -1
	}

	return grid[rowIndex][colIndex]
}

func isTreeVisible(treeRow, treeCol, otherRow, otherCol int, grid [][]int) bool {
	treeHeight := getTreeHeightAt(treeRow, treeCol, grid)
	otherTreeHeight := getTreeHeightAt(otherRow, otherCol, grid)

	// a tree against an edge is always visible
	if otherTreeHeight == -1 {
		return true
	}

	if otherTreeHeight >= treeHeight {
		return false
	}

	// left
	if otherCol < treeCol {
		return isTreeVisible(treeRow, treeCol, otherRow, otherCol-1, grid)
	}
	// right
	if otherCol > treeCol {
		return isTreeVisible(treeRow, treeCol, otherRow, otherCol+1, grid)
	}
	// down
	if otherRow > treeRow {
		return isTreeVisible(treeRow, treeCol, otherRow+1, otherCol, grid)
	}

	// up
	return isTreeVisible(treeRow, treeCol, otherRow-1, otherCol, grid)

}

func findVisibleTrees(grid [][]int) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			visibleFromLeft := isTreeVisible(i, j, i, j-1, grid)
			visibleFromRight := isTreeVisible(i, j, i, j+1, grid)
			visibleFromTop := isTreeVisible(i, j, i-1, j, grid)
			visibleFromBottom := isTreeVisible(i, j, i+1, j, grid)
			if visibleFromRight || visibleFromTop || visibleFromLeft || visibleFromBottom {
				count++
			}
		}
	}
	return count
}

func getViewingDistance(fromRow, fromCol, toRow, toCol int, grid [][]int) int {
	treeHeight := getTreeHeightAt(fromRow, fromCol, grid)
	otherTreeHeight := getTreeHeightAt(toRow, toCol, grid)
	count := 1
	// this is an edge
	if otherTreeHeight == -1 {
		return 0
	}
	if otherTreeHeight >= treeHeight {
		return count
	}

	// left
	if toCol < fromCol {
		return count + getViewingDistance(fromRow, fromCol, toRow, toCol-1, grid)
	}
	// right
	if toCol > fromCol {
		return count + getViewingDistance(fromRow, fromCol, toRow, toCol+1, grid)
	}
	// down
	if toRow > fromRow {
		return count + getViewingDistance(fromRow, fromCol, toRow+1, toCol, grid)
	}

	// up
	return count + getViewingDistance(fromRow, fromCol, toRow-1, toCol, grid)

}

func getHighestScenicScore(grid [][]int) int {
	var max int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			left := getViewingDistance(i, j, i, j-1, grid)
			right := getViewingDistance(i, j, i, j+1, grid)
			top := getViewingDistance(i, j, i-1, j, grid)
			bottom := getViewingDistance(i, j, i+1, j, grid)
			total := left * right * top * bottom
			max = int(math.Max(float64(max), float64(total)))
		}
	}

	return max

}

func main() {
	args := os.Args[1:]
	grid := processInput(args[0])
	numVisibleTrees := findVisibleTrees(grid)
	fmt.Println(numVisibleTrees)

	fmt.Println(getHighestScenicScore(grid))
}
