package aocutils

import (
	"log"
	"os"
	"strings"
)

func GetRawInput() string {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal("problem reading input")
	}

	return string(data)
}

func GetInputRows() []string {
	input := GetRawInput()
	return strings.Split(input, "\n")
}

func GetIntMatrixInput(separator string) [][]int {
	input := GetInputRows()
	matrix := make([][]int, len(input))

	for y, str := range input {
		nums := strings.Split(str, separator)
		matrix[y] = make([]int, len(nums))

		for x, n := range nums {
			matrix[y][x] = Stoi(n)
		}
	}

	return matrix
}

func GetSingleRowInputs(separator string) []string {
	input := GetRawInput()
	return strings.Split(input, separator)
}
