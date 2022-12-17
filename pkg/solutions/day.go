package solutions

type Day interface {
	DayNumber() int
	ResultsA() (example int, result int)
	SolutionA(input []string) int
	ResultsB() (example int, result int)
	SolutionB(input []string) int
}
