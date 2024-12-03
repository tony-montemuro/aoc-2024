package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func scanNumber(input string, index *int, stopChar byte) int {
	number, i := "", 1

	for i <= 3 && input[*index] != stopChar {
		if unicode.IsDigit(rune(input[*index])) {
			number += string(input[*index])
		} else {
			return -1
		}
		*index++
	}

	if input[*index] != stopChar {
		return -1
	}
	return aocutils.Stoi(number)
}

func truncateInput(inputPtr *string, indexPtr *int) {
	input := *inputPtr
	*inputPtr = input[*indexPtr:]
	*indexPtr = strings.Index(*inputPtr, "mul(")
}

func scan(input string) int {
	output := 0
	index := strings.Index(input, "mul(")

	for index != -1 {
		index += 4
		n1 := scanNumber(input, &index, ',')
		if n1 == -1 {
			truncateInput(&input, &index)
			continue
		}
		index++
		n2 := scanNumber(input, &index, ')')
		if n2 == -1 {
			truncateInput(&input, &index)
			continue
		}

		output += n1 * n2
		truncateInput(&input, &index)
	}
	return output
}

func getEnabledSections(input string) string {
	output := ""
	do, dont := "do()", "don't()"
	doIndex := strings.Index(input, do)
	dontIndex := strings.Index(input, dont)
	isEnabled := true

	for doIndex != -1 && dontIndex != -1 {
		smallerIndex, function := doIndex, do
		if dontIndex < doIndex {
			smallerIndex = dontIndex
			function = dont
		}

		if isEnabled {
			output += input[0:smallerIndex]
		}

		isEnabled = doIndex < dontIndex
		input = input[smallerIndex+len(function):]
		doIndex = strings.Index(input, do)
		dontIndex = strings.Index(input, dont)
	}

	if isEnabled {
		if dontIndex != -1 {
			output += input[0:dontIndex]
		} else {
			output += input
		}
	} else if doIndex != -1 {
		output += input[doIndex+len(do):]
	}

	return output
}

func main() {
	input := aocutils.GetRawInput()
	fmt.Println(scan(input))
	fmt.Println(scan(getEnabledSections(input)))
}
