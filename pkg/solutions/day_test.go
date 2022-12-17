package solutions_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"simonkarman.nl/advent-of-code-2022/pkg/solutions"
	"simonkarman.nl/advent-of-code-2022/pkg/utils"
	"testing"
)

func TestSolutions(t *testing.T) {
	days := []solutions.Day{
		solutions.Day00{},
		solutions.Day01{},
		solutions.Day02{},
	}
	for _, day := range days {
		number := day.DayNumber()
		t.Run(fmt.Sprintf("Day%02d", number), func(t *testing.T) {
			var exampleInput, input = utils.GetInputs(number)

			// A
			exampleA, resultA := day.ResultsA()
			assert.Equal(t, exampleA, day.SolutionA(exampleInput))
			assert.Equal(t, resultA, day.SolutionA(input))

			// B
			exampleB, resultB := day.ResultsB()
			assert.Equal(t, exampleB, day.SolutionB(exampleInput))
			assert.Equal(t, resultB, day.SolutionB(input))
		})
	}
}
