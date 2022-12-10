package main

import (
	"adventOfCode2022/pkg/ds/list"
	"adventOfCode2022/pkg/util"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	UP    = Direction("U")
	Down  = Direction("D")
	Left  = Direction("L")
	Right = Direction("R")
)

func moveTo(current []int, direction Direction, numMoves int) []int {
	cx := current[0]
	cy := current[1]

	switch direction {
	case Left:
		return []int{cx - numMoves, cy}
	case Right:
		return []int{cx + numMoves, cy}
	case UP:
		return []int{cx, cy + numMoves}
	case Down:
		return []int{cx, cy - numMoves}
	default:
		return []int{cx, cy}
	}
}

func distance(x1, y1, x2, y2 int) int {
	asq := math.Pow(float64(x1-x2), 2)
	bsq := math.Pow(float64(y1-y2), 2)
	c := math.Sqrt(asq + bsq)

	return int(c)
}

func oppositeDirection(dir Direction) (Direction, error) {
	switch dir {
	case Left:
		return Right, nil
	case Right:
		return Left, nil
	case UP:
		return Down, nil
	case Down:
		return UP, nil
	default:
		return UP, errors.New("unsupported direction")
	}
}

func getKey(x, y int) string {
	return "(" + fmt.Sprintf("%d", x) + "," + fmt.Sprintf("%d", y) + ")"
}

func initCoords(size int) *list.List[[]int] {
	coords := list.New[[]int]()
	for i := 0; i < size; i++ {
		coords.Push([]int{0, 0})
	}

	return coords
}

func nextCoords(headX, headY, tailX, tailY int) (x, y int) {
	// tail coordinates do not change if distance <= 1 since
	// they are touching already
	if distance(headX, headY, tailX, tailY) <= 1 {
		return tailX, tailY
	}

	// same row
	if headY == tailY {
		if headX > tailX {
			return headX - 1, tailY
		}
		return headX + 1, tailY
	}
	// same column
	if headX == tailX {
		if headY > tailY {
			return tailX, headY - 1
		}
		return tailX, headY + 1
	}

	// diagonal
	if headX > tailX && headY > tailY {
		return tailX + 1, tailY + 1
	}
	if headY > tailY {
		return tailX - 1, tailY + 1
	}
	if headX < tailX {
		return tailX - 1, tailY - 1
	}
	return tailX + 1, tailY - 1
}

func main() {
	args := os.Args[1:]
	coords := initCoords(10)
	tail := coords.TailNode()
	keys := make(map[string]bool)

	for line := range util.ReadInputFileLines(args[0]) {
		split := strings.Split(line, " ")
		dir := split[0]
		num, _ := strconv.ParseInt(split[1], 10, 64)

		head := coords.HeadNode()
		for i := 0; i < int(num); i++ {
			head.Set(moveTo(head.Value(), Direction(dir), 1))
			currNode := head
			for currNode.Next() != nil && distance(currNode.Value()[0], currNode.Value()[1], currNode.Next().Value()[0], currNode.Next().Value()[1]) > 1 {
				x, y := nextCoords(currNode.Value()[0], currNode.Value()[1], currNode.Next().Value()[0], currNode.Next().Value()[1])
				currNode = currNode.Next()
				currNode.Set([]int{x, y})
			}
			keys[getKey(tail.Value()[0], tail.Value()[1])] = true
			fmt.Printf("Tail Coords %s\n", getKey(tail.Value()[0], tail.Value()[1]))
		}
		fmt.Printf("After %s: %s\n", line, coords)
	}

	fmt.Println(len(keys))
}
