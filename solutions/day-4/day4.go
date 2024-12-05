package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func getPaddedInput(w, h int, input []string) []string {
	paddedInput := []string{}
	for _, row := range input {
		pad := strings.Repeat(".", w)
		paddedInput = append(paddedInput, pad+row+pad)
	}

	paddedRow := strings.Repeat(".", (2*w)+len(input[0]))
	for i := 0; i < h; i++ {
		paddedInput = slices.Insert(paddedInput, 0, paddedRow)
		paddedInput = append(paddedInput, paddedRow)
	}

	return paddedInput
}

func getTargetCountPerPosition(x, y int, target string, input []string) int {
	targetCount := 0
	dirs := []int{-1, 0, 1}

	for _, yDir := range dirs {
		for _, xDir := range dirs {
			word := string(input[y][x])
			i := 1
			for len(word) < len(target) {
				word += string(input[y+(yDir*i)][x+(xDir*i)])
				i++
			}
			if word == target {
				targetCount++
			}
		}
	}

	return targetCount
}

func reverse(s string) string {
	bytes := make([]byte, len(s))
	i, j := 0, len(s)-1

	for i <= j {
		bytes[i], bytes[j] = s[j], s[i]
		i++
		j--
	}

	return string(bytes)
}

func isXPosition(x, y int, target string, input []string) bool {
	reversedTarget := reverse(target)
	w1 := string(input[y-1][x-1]) + string(input[y][x]) + string(input[y+1][x+1])
	w2 := string(input[y+1][x-1]) + string(input[y][x]) + string(input[y-1][x+1])
	return (w1 == target || w1 == reversedTarget) && (w2 == target || w2 == reversedTarget)
}

func main() {
	input := aocutils.GetInputRows()
	target := "XMAS"
	padLength := len(target) - 1
	input = getPaddedInput(padLength, padLength, input)
	p1, p2 := 0, 0

	for y := padLength; y < len(input)-padLength; y++ {
		for x := padLength; x < len(input[y])-padLength; x++ {
			p1 += getTargetCountPerPosition(x, y, target, input)
			if isXPosition(x, y, target[1:], input) {
				p2++
			}
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
