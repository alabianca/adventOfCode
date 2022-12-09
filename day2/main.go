package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const Rock = 'A'
const Paper = 'B'
const Scissors = 'C'

//const RockMe = 'X'
//const PaperMe = 'Y'
//const ScissorsMe = 'Z'
const Lose = 'X'
const Draw = 'Y'
const Win = 'Z'

func readInput(path string) <-chan []byte {
	output := make(chan []byte)

	go func() {
		defer close(output)

		file, err := os.Open(path)
		if err != nil {
			log.Printf("Error %s\n", err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			b := []byte(scanner.Text())
			// drop the space
			output <- []byte{b[0], b[2]}
		}
	}()

	return output
}

func getPointsForSelection(selection byte) int {
	return int(selection - byte(Rock) + 1)
}

func calculatePoints(opponent, targetOutcome byte) int {
	var outcome int

	mySelection := byte(Rock)
	switch targetOutcome {
	case Win:
		if opponent == Rock {
			mySelection = Paper
		} else if opponent == Scissors {
			mySelection = Rock
		} else {
			mySelection = Scissors
		}
		outcome = 6

	case Draw:
		mySelection = opponent
		outcome = 3
	case Lose:
		if opponent == Rock {
			mySelection = Scissors
		} else if opponent == Scissors {
			mySelection = Paper
		} else {
			mySelection = Rock
		}
		outcome = 0
	}

	return outcome + getPointsForSelection(mySelection)
}

func main() {
	args := os.Args[1:]
	var total int64

	for b := range readInput(args[0]) {
		total += int64(calculatePoints(b[0], b[1]))
	}

	fmt.Printf("Total is %d\n", total)
}
