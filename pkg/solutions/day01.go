package solutions

import (
	"fmt"
	"sort"
	"strconv"
)

func SolutionDay01A(input []string) int {
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

func SolutionDay01B(input []string) int {
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
