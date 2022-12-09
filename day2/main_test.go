package main

import "testing"

func Test_CalculatePoints(t *testing.T) {
	total := calculatePoints(Rock, Draw) + calculatePoints(Paper, Lose) + calculatePoints(Scissors, Win)
	if total != 12 {
		t.Errorf("Expected %d, but got %d\n", 15, total)
	}
}
