# Advent of Code 2022
This repository contains the solutions to Advent of Code 2022 by [Simon Karman](https://www.simonkarman.nl). [Advent of Code](https://adventofcode.com/2022/leaderboard/private/view/718869) is an Advent calendar of small programming puzzles for a variety of skill sets and skill levels that can be solved in any programming language you like.

All the solutions in this repository are written in GoLang.

> Disclaimer: This repository contains the answer to puzzles of Advent of Code 2022. If you don't want spoilers, then please don't look any further.

# Getting Started
To get started download all dependencies.
```bash
go mod download
```

To verify the solutions of each day you can run the tests.
```bash
go test pkg/solutions/day_test.go -v
```

# Structure
The section below explains the structure of the files in this project.

## Days
The `input/` directory contains two input files for each day. One for the example input (`day01-example.txt`) and one for the answer input (`day01.txt`).

The solutions for each day can be found in the `pkg/solutions/` directory. The solutions of a day are stored in the corresponding file of that day (`day01.go`, `day02.go`, ect). Each day exports a struct that conforms to the interface `Day`, which can be found in `day.go`.

> Note: To easily add a new day you can copy `pkg/solutions/day00.go`, `input/day00.txt`, and `input/day00-example.txt`.

## Tests
All days are unit tested. You can run smaller s of the tests.
```bash
# run all tests of all days
go test pkg/solutions/day_test.go

# run all tests of single day
go test pkg/solutions/day_test.go -run 'TestSolutions/Day01'
```
