package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func main() {
	input := aocutils.GetRawInput()
	left, right := []int{}, []int{}
	leftCounts, rightCounts := make(map[int]int), make(map[int]int)
	sum, similarityScore := 0, 0

	for _, row := range strings.Split(string(input), "\n") {
		nums := strings.Split(row, "   ")
		l, r := aocutils.Stoi(nums[0]), aocutils.Stoi(nums[1])

		left = append(left, l)
		right = append(right, r)

		leftCounts[l]++
		rightCounts[r]++
	}

	slices.SortFunc(left, func(a, b int) int { return a - b })
	slices.SortFunc(right, func(a, b int) int { return a - b })

	for i := 0; i < len(left); i++ {
		sum += aocutils.AbsInt(right[i] - left[i])
	}

	for key, val := range rightCounts {
		similarityScore += key * val * leftCounts[key]
	}

	fmt.Println(sum)
	fmt.Println(similarityScore)
}
