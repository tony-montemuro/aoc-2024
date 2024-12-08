package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

type operation func(int, int) int

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func concat(x, y int) int {
	exp := 0
	right := y
	for right > 0 {
		right /= 10
		exp++
	}
	return int(math.Pow(10, float64(exp)))*x + y
}

func isValidEquation(nums []int, target int, ops ...operation) bool {
	if len(nums) == 2 {
		for _, op := range ops {
			if op(nums[1], nums[0]) == target {
				return true
			}
		}
		return false
	}

	first, second := nums[len(nums)-1], nums[len(nums)-2]
	for _, op := range ops {
		if isValidEquation(append(append([]int{}, nums[:len(nums)-2]...), op(first, second)), target, ops...) {
			return true
		}
	}
	return false
}

func main() {
	input := aocutils.GetInputRows()
	p1, p2 := 0, 0

	for _, row := range input {
		split := strings.Split(row, ": ")
		target := aocutils.Stoi(split[0])

		strNums, nums := strings.Split(split[1], " "), []int{}
		for i := len(strNums) - 1; i >= 0; i-- {
			nums = append(nums, aocutils.Stoi(strNums[i]))
		}

		if isValidEquation(nums, target, add, mul) {
			p1 += target
			p2 += target
		} else if isValidEquation(nums, target, add, mul, concat) {
			p2 += target
		}
	}
	fmt.Println(p1, p2)
}
