package solutions

import (
	"fmt"
	"sort"
	"strconv"
)

type Day01 struct{}

func (d Day01) DayNumber() int {
	return 1
}

func (d Day01) SolutionA(input []string) int {
	current, max := 0, 0
	for i := 0; i < len(input); i++ {
		calories, err := strconv.Atoi(input[i])
		if err != nil {
			current = 0
		} else {
			current += calories
			if current > max {
				max = current
			}
		}
	}
	return max
}

func (d Day01) ResultsA() (int, int) {
	return 24000, 69836
}

func (d Day01) SolutionB(input []string) int {
	elves := []int{0}
	for i := 0; i < len(input); i++ {
		calories, err := strconv.Atoi(input[i])
		if err != nil {
			elves = append(elves, 0)
		} else {
			elves[len(elves)-1] += calories
		}
	}
	fmt.Println(len(elves))
	sort.Ints(elves)
	sum := 0
	for i := len(elves) - 3; i < len(elves); i++ {
		sum += elves[i]
	}
	return sum
}

func (d Day01) ResultsB() (int, int) {
	return 45000, 207968
}
