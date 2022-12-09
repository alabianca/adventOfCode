package main

import (
	"adventOfCode2022/pkg/ds/queue"
	"adventOfCode2022/pkg/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// width of column = 4
// distance from start of column X to column Y = (Y - X)*columnWidth + (Y-X)
// number of columns in line = (len(line)+1-(len(line)+1%columnWidth)/columnWidth
func getNumOfColumnsInLine(line string, columnWidth int) int {
	gutterOffset := 1
	length := len(line)
	withOffset := gutterOffset + length

	return (withOffset - (withOffset % columnWidth)) / columnWidth
}

func processInput(path string) (<-chan []rune, <-chan []int, <-chan struct{}) {
	stackChan := make(chan []rune)
	instructionChan := make(chan []int)
	done := make(chan struct{})

	go func() {
		defer close(stackChan)
		defer close(instructionChan)
		defer close(done)

		instructionStartToken := "move"
		stackLineEndToken := "]"
		columnWidth := 4
		for line := range util.ReadInputFileLines(path) {
			// is this an instruction or are we still parsing the stack illustration?
			if strings.HasPrefix(line, instructionStartToken) {
				split := strings.Split(line, " ")
				// only send valid instructions ["move", "5", "from", "4", "to", "9"]
				if len(split) == 6 {
					count, _ := strconv.ParseInt(split[1], 10, 64)
					from, _ := strconv.ParseInt(split[3], 10, 64)
					to, _ := strconv.ParseInt(split[5], 10, 64)
					instructionChan <- []int{int(count), int(from), int(to)}

				}
			}
			// is this a stack line?
			if strings.HasSuffix(strings.TrimSpace(line), stackLineEndToken) {
				numCols := getNumOfColumnsInLine(line, columnWidth)
				columns := make([]rune, numCols)
				for i := 0; i < len(columns); i++ {
					endIndex := (i + 1) * columnWidth
					startIndex := i * columnWidth
					// edge case at the end
					var col string
					if i == len(columns)-1 {
						col = strings.TrimSpace(line[startIndex:])
					} else {
						col = strings.TrimSpace(line[startIndex:endIndex])
					}

					if len(col) > 0 {
						columns[i] = rune(col[1])
					}
				}
				cp := make([]rune, numCols)
				copy(cp, columns)
				stackChan <- cp
			}
			// allocate more space in the stacks slice if needed
			//diff := len(stacks) - numCols
		}

		// end of file
		done <- struct{}{}

	}()

	return stackChan, instructionChan, done
}

func moveFromToQueue(count int, from *queue.Queue[rune], to *queue.Queue[rune]) error {
	spliced := from.DequeueRange(from.Size()-count, from.Size())
	to.EnqueueBackList(spliced)
	fmt.Printf("Dequeued %s\n", spliced)
	//for i := 0; i < count; i++ {
	//	v, err := from.DeQueueBack()
	//	if err != nil {
	//		return err
	//	}
	//
	//	to.EnqueueBack(v)
	//}

	return nil
}

func main() {
	args := os.Args[1:]

	stackValues, instructions, done := processInput(args[0])
	queues := make([]*queue.Queue[rune], 0)

	exit := false
	for !exit {
		select {
		case values := <-stackValues:
			fmt.Println(values)
			// how many more queues do I need to add?
			diff := len(values) - len(queues)
			queues = append(queues, make([]*queue.Queue[rune], diff)...)
			// enqueue all the values into the appropriate queue
			for i, val := range values {
				if queues[i] == nil {
					queues[i] = queue.New[rune]()
				}
				if val > 0 {
					queues[i].EnqueueFront(val)
				}
				//fmt.Printf("Queue %d : %s\n", i, queues[i])

			}
		case instruction := <-instructions:
			count := instruction[0]
			from := instruction[1]
			to := instruction[2]
			fmt.Printf("Moving %d values from %s to %s\n", count, queues[from-1], queues[to-1])
			moveFromToQueue(count, queues[from-1], queues[to-1])
			fmt.Printf("After Move %s => %s\n", queues[from-1], queues[to-1])
		case <-done:
			exit = true
		}
	}

	message := strings.Builder{}
	for i := 0; i < len(queues); i++ {
		if c, ok := queues[i].PeekBack(); ok {
			message.WriteRune(c)
		}
	}

	fmt.Printf("%s\n", message.String())
}
