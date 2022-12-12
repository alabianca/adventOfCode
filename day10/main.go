package main

import (
	"adventOfCode2022/pkg/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction string

const (
	addInstruction  = Instruction("addx")
	noopInstruction = Instruction("noop")
)

func makeScreen() string {
	var out string
	for i := 0; i < 240; i++ {
		out += "."
	}
	return out
}

func main() {
	args := os.Args[1:]

	regCycleVals := make([]int, 0)
	regCycleVals = append(regCycleVals, 1)
	screen := makeScreen()
	
	for line := range util.ReadInputFileLines(args[0]) {
		split := strings.Split(line, " ")
		instruction := Instruction(split[0])
		var crtIndex int
		switch instruction {
		case addInstruction:
			x, _ := strconv.ParseInt(split[1], 10, 64)
			prev := regCycleVals[len(regCycleVals)-1]
			regCycleVals = append(regCycleVals, []int{prev, prev + int(x)}...)
			crtIndex = regCycleVals[len(regCycleVals)-1]

		case noopInstruction:
			prev := regCycleVals[len(regCycleVals)-1]
			regCycleVals = append(regCycleVals, prev)
		}
	}

	var total int
	for i := 19; i < len(regCycleVals); i = i + 40 {
		total = total + regCycleVals[i]*(i+1)
	}

	//screen := make([]string, len(regCycleVals))
	//for i := 0; i < len(regCycleVals); i++ {
	//	middle := regCycleVals[i]
	//	left := regCycleVals[i] - 1
	//	right := regCycleVals[i] + 1
	//	screen[left] = "#"
	//	screen[right] = "#"
	//	screen[middle] = "#"
	//}
	//
	//// render screen
	//for i := 0; i < len(screen); i++ {
	//	if len(screen[i]) == 0 {
	//		fmt.Printf(".")
	//	} else {
	//		fmt.Printf("%s\n", screen[i])
	//	}
	//	if (i+1)%40 == 0 {
	//		fmt.Println()
	//	}
	//}

	fmt.Println(total)
}
