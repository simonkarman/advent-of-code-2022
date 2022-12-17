package solutions

import (
	"strings"
)

type Day03 struct{}

func (d Day03) DayNumber() int {
	return 3
}

func toPriority(item int) int {
	if item > 96 {
		return int(item) - 96
	} else {
		return int(item) - 64 + 26
	}
}

func (d Day03) SolutionA(input []string) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		rucksack := input[i]
		halfSize := len(rucksack) / 2
		second := rucksack[halfSize:]
		for x := 0; x < halfSize; x++ {
			item := rucksack[x]
			if strings.Index(second, string(item)) != -1 {
				sum += toPriority(int(item))
				break
			}
		}
	}
	return sum
}

func (d Day03) ResultsA() (int, int) {
	return 157, 7746
}

func (d Day03) SolutionB(input []string) int {
	sum := 0
	for i := 0; i < len(input); i += 3 {
		rucksackA := input[i]
		rucksackB := input[i+1]
		rucksackC := input[i+2]
		for x := 0; x < len(rucksackA); x++ {
			item := rucksackA[x]
			isInB := strings.Index(rucksackB, string(item)) != -1
			isInC := strings.Index(rucksackC, string(item)) != -1
			if isInB && isInC {
				sum += toPriority(int(item))
				break
			}
		}
	}
	return sum
}

func (d Day03) ResultsB() (int, int) {
	return 70, 2604
}
