package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func getEquation(machine string) ([]int, []int) {
	lhs, rhs := make([]int, 4), make([]int, 2)
	lines := strings.Split(machine, "\n")

	for i, line := range lines {
		s := strings.Split(strings.Split(line, ": ")[1], ", ")
		x, y := aocutils.Stoi(s[0][2:]), aocutils.Stoi(s[1][2:])
		if i < len(lines)-1 {
			lhs[i] = x
			lhs[i+2] = y
		} else {
			rhs[0] = x
			rhs[1] = y
		}
	}

	return lhs, rhs
}

func getDet(lhs []int) int {
	return lhs[0]*lhs[3] - lhs[1]*lhs[2]
}

func getButtonPresses(lhs []int, rhs []int) (int, int) {
	a := lhs[0]*rhs[0] + lhs[1]*rhs[1]
	b := lhs[2]*rhs[0] + lhs[3]*rhs[1]
	return a, b
}

func getMinTokens(lhs []int, rhs []int, maxButtonPress int) int {
	det := getDet(lhs)
	lhs[0], lhs[3] = lhs[3], lhs[0]
	lhs[1], lhs[2] = -lhs[1], -lhs[2]
	a, b := getButtonPresses(lhs, rhs)

	if a%det != 0 || b%det != 0 {
		return 0
	}
	a /= det
	b /= det

	if a > maxButtonPress || b > maxButtonPress {
		return 0
	}
	tokens := 3*a + b
	return tokens
}

func main() {
	input := aocutils.GetRawInput()
	machines := strings.Split(input, "\n\n")
	p1, p2 := 0, 0
	for _, machine := range machines {
		// part 1
		lhs, rhs := getEquation(machine)
		p1 += getMinTokens(lhs, rhs, 100)

		// part 2
		lhs, rhs = getEquation(machine)
		rhs[0] += 10000000000000
		rhs[1] += 10000000000000
		p2 += getMinTokens(lhs, rhs, math.MaxInt)
	}
	fmt.Println(p1, p2)
}
