package solutions_test

import (
	"github.com/stretchr/testify/assert"
	"simonkarman.nl/advent-of-code-2022/pkg/solutions"
	"simonkarman.nl/advent-of-code-2022/pkg/utils"
	"testing"
)

var exampleInput02, input02 = utils.GetInputs(2)

func TestSolutionDay02A(t *testing.T) {
	assert.Equal(t, 15, solutions.SolutionDay02A(exampleInput02))
	assert.Equal(t, 11150, solutions.SolutionDay02A(input02))
}

func TestSolutionDay02B(t *testing.T) {
	assert.Equal(t, 12, solutions.SolutionDay02B(exampleInput02))
	assert.Equal(t, 8295, solutions.SolutionDay02B(input02))
}
