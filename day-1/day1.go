package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func strToInt(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		log.Fatalf("problem converting %s to int", s)
	}

	return n
}

func ascSort(a, b int) int {
	return a - b
}

func countInList(intMap *map[int]int, l []int) {
	m := *intMap
	for _, n := range l {
		m[n]++
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("problem reading input")
	}

	l1, l2 := []int{}, []int{}
	for _, row := range strings.Split(string(input), "\n") {
		nums := strings.Split(row, "   ")
		n1, n2 := strToInt(nums[0]), strToInt(nums[1])
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}
	slices.SortFunc(l1, ascSort)
	slices.SortFunc(l2, ascSort)

	sum := 0
	for i := 0; i < len(l1); i++ {
		diff := l2[i] - l1[i]

		// abs value
		if diff > 0 {
			sum += diff
		} else {
			sum -= diff
		}
	}

	leftCounts := make(map[int]int)
	rightCounts := make(map[int]int)
	for _, n := range l1 {
		leftCounts[n] = 0
		leftCounts[n] = 0
	}

	countInList(&leftCounts, l1)
	countInList(&rightCounts, l2)

	similarityScore := 0
	for key, val := range rightCounts {
		similarityScore += key * val * leftCounts[key]
	}

	fmt.Println(sum)
	fmt.Println(similarityScore)
}
