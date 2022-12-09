package util

import (
	"bufio"
	"log"
	"os"
)

func ReadInputFileLines(path string) <-chan string {
	output := make(chan string)

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
			output <- scanner.Text()
		}
	}()

	return output
}
