package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	rootFolder = filepath.Join(filepath.Dir(b), "../..")
)

func getInput(fileName string) []string {
	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	// Read all lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return lines
}

func GetInputs(day int) ([]string /*example*/, []string /*real*/) {
	exampleFile := fmt.Sprintf("%s/input/day%02d-example.txt", rootFolder, day)
	inputFile := fmt.Sprintf("%s/input/day%02d.txt", rootFolder, day)
	return getInput(exampleFile), getInput(inputFile)
}
