package main

import (
	"adventOfCode2022/pkg/ds/queue"
	"adventOfCode2022/pkg/util"
	"fmt"
	"math"
	"os"
	"strings"
)

func parseInput(path string) ([][]int, int, int, int, int) {
	rowIndex := 0
	out := make([][]int, 0)
	startX := 0
	startY := 0
	endX := 0
	endY := 0
	for line := range util.ReadInputFileLines(path) {
		row := make([]int, len(line))
		for i, c := range strings.Split(line, "") {
			if c == "S" {
				startX = i
				startY = rowIndex
				row[i] = 0
			} else if c == "E" {
				endX = i
				endY = rowIndex
				row[i] = 25
			} else {
				row[i] = int([]byte(c)[0]) - 'a'
			}

		}
		rowIndex++
		out = append(out, row)
	}

	return out, startX, startY, endX, endY
}

func getAdjMatrix(inMatrix [][]int) [][]int {
	adjMatrix := make([][]int, len(inMatrix)*len(inMatrix[0]))
	// init with 0 values
	for i, _ := range adjMatrix {
		adjMatrix[i] = make([]int, len(inMatrix)*len(inMatrix[0]))
	}

	nodesConnect := func(x1, y1, x2, y2 int) bool {
		src := inMatrix[x1][y1]
		target := inMatrix[x2][y2]
		return target <= src || target-src <= 1
	}

	sourceIndex := 0
	for i, row := range inMatrix {
		for j, _ := range row {
			// left
			if j-1 >= 0 && nodesConnect(i, j, i, j-1) {
				adjMatrix[sourceIndex][sourceIndex-1] = 1
			}
			// right
			if j+1 < len(row) && nodesConnect(i, j, i, j+1) {
				adjMatrix[sourceIndex][sourceIndex+1] = 1
			}
			// up
			if i-1 >= 0 && nodesConnect(i, j, i-1, j) {
				adjMatrix[sourceIndex][sourceIndex-len(row)] = 1
			}
			// down
			if i+1 < len(inMatrix) && nodesConnect(i, j, i+1, j) {
				adjMatrix[sourceIndex][sourceIndex+len(row)] = 1
			}

			sourceIndex++
		}
	}

	return adjMatrix
}

func prettyPrintAdjMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		str := ""
		for j := 0; j < len(matrix[i]); j++ {
			str += fmt.Sprintf("%d ", matrix[i][j])
		}
		fmt.Println(str)
	}
}

func bfs(graph [][]int, start, end int) []int {
	numNodes := len(graph)
	visited := make([]bool, numNodes)
	prev := make([]int, numNodes)
	// init prev
	for i := range prev {
		prev[i] = -1
	}

	q := queue.New[int]()
	q.EnqueueBack(start)

	for {
		v, err := q.DeQueueFront()
		if err != nil || v == end {
			break
		}
		visited[v] = true
		neighbors := graph[v]
		for i, n := range neighbors {
			if !visited[i] && n == 1 {
				q.EnqueueBack(i)
				prev[i] = v
			}
		}
	}

	return prev
}

func shortestPath(prevList []int, endNode int) []int {
	path := make([]int, 0)
	for i := endNode; i != -1; i = prevList[i] {
		path = append(path, i)
	}

	return path
}

func getMinDistance(dist []float64, sptSet []bool) int {
	min := math.Inf(1)
	minIndex := -1

	for i := 0; i < len(dist); i++ {
		if dist[i] != -1 && dist[i] < min && !sptSet[i] {
			min = dist[i]
			minIndex = i
		}
	}

	return minIndex
}

func dijkstra(graph [][]int, src, target int) []float64 {
	numVertices := len(graph)
	sptSet := make([]bool, numVertices)
	dist := make([]float64, numVertices)
	// init all distances to INF and init sptSet to false for every vertex
	for i := 0; i < numVertices; i++ {
		dist[i] = math.Inf(1)
		sptSet[i] = false
	}

	// distance from source to source is 0
	dist[src] = 0

	// get distance from source to every node
	for i := 0; i < numVertices; i++ {
		minIndex := getMinDistance(dist, sptSet)
		sptSet[minIndex] = true
		for j := 0; j < numVertices; j++ {
			if !sptSet[j] && graph[minIndex][j] != 0 && dist[minIndex] < math.Inf(1) && dist[minIndex]+float64(graph[j][minIndex]) < dist[j] {
				dist[j] = dist[minIndex] + float64(graph[j][minIndex])
			}
		}
	}

	return dist
}

func main() {
	args := os.Args[1:]
	matrix, startX, startY, endX, endY := parseInput(args[0])
	adjMatrix := getAdjMatrix(matrix) // n * n matrix
	//prettyPrintAdjMatrix(adjMatrix)
	startNode := (len(matrix[0]) * startY) + startX
	endNode := (len(matrix[0]) * endY) + endX
	dist := dijkstra(adjMatrix, startNode, endNode)
	//path := bfs(adjMatrix, startNode, endNode)
	//fmt.Println(shortestPath(path, endNode))
	//
	//fmt.Println(path)
	//fmt.Println(dist)
	fmt.Printf("Min Distance to %d is %d\n", endNode, dist[endNode])
}
