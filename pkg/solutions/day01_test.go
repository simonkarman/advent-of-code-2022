package solutions_test

import (
	"github.com/stretchr/testify/assert"
	"simonkarman.nl/advent-of-code-2022/pkg/solutions"
	"simonkarman.nl/advent-of-code-2022/pkg/utils"
	"testing"
)

var exampleInput01, input01 = utils.GetInputs(1)

func TestSolutionDay01A(t *testing.T) {
	assert.Equal(t, 24000, solutions.SolutionDay01A(exampleInput01))
	assert.Equal(t, 69836, solutions.SolutionDay01A(input01))
}

func TestSolutionDay01B(t *testing.T) {
	assert.Equal(t, 45000, solutions.SolutionDay01B(exampleInput01))
	assert.Equal(t, 207968, solutions.SolutionDay01B(input01))
}
