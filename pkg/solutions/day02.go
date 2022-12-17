package solutions

import (
	"fmt"
	"strings"
)

type Day02 struct{}

func (d Day02) DayNumber() int {
	return 2
}

// 1 Rock
// 2 Paper
// 3 Scissors
func toShape(input string, sheet string) int {
	index := strings.Index(sheet, input)
	if index == -1 {
		panic(fmt.Sprintf("shape: cannot find '%s' in %s", input, sheet))
	}
	return index + 1
}

func toScore(opponent, me int) int {
	const (
		Rock     int = 1
		Paper        = 2
		Scissors     = 3
	)
	// draw
	if opponent == me {
		return 3
	}
	// 21 - lose
	// 13 - lose
	// 32 - lose
	if (opponent == Paper && me == Rock) || (opponent == Rock && me == Scissors) || (opponent == Scissors && me == Paper) {
		return 0
	}
	// 12 - win
	// 31 - win
	// 23 - win
	return 6
}

func (d Day02) SolutionA(input []string) int {
	score := 0
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			continue
		}
		round := strings.Split(input[i], " ")
		opponent := toShape(round[0], "ABC")
		me := toShape(round[1], "XYZ")
		score = score + me + toScore(opponent, me)
	}
	return score
}

func (d Day02) ResultsA() (int, int) {
	return 15, 11150
}

func (d Day02) SolutionB(input []string) int {
	score := 0
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			continue
		}
		round := strings.Split(input[i], " ")
		opponent := toShape(round[0], "ABC")
		targetScore := (toShape(round[1], "XYZ") - 1) * 3
		for me := 1; me <= 3; me++ {
			hypotheticalScore := toScore(opponent, me)
			if hypotheticalScore == targetScore {
				score = score + me + targetScore
				break
			}
		}
	}
	return score
}

func (d Day02) ResultsB() (int, int) {
	return 12, 8295
}
