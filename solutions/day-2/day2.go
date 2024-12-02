package main

import (
	"fmt"
	"strings"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func isReportSafe(originalReport []int, skipLevel int) bool {
	prev := 0
	report := []int{}

	// build report, skipping a level if specified
	for i, level := range originalReport {
		if i != skipLevel {
			report = append(report, level)
		}
	}

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		absDiff := aocutils.AbsInt(diff)

		// if difference is less than 1, greater than 3, or sign is different, report is unsafe
		if absDiff < 1 ||
			absDiff > 3 ||
			(i > 1 && (prev < 0 && diff > 0) || (prev > 0 && diff < 0)) {
			return false
		}

		prev = diff
	}

	return true
}

func isToleratedReportSafe(report []int) bool {
	for i := range len(report) {
		if isReportSafe(report, i) {
			return true
		}
	}
	return false
}

func main() {
	input := aocutils.GetRawInput()
	p1, p2 := 0, 0

	for _, row := range strings.Split(input, "\n") {
		nums := strings.Split(row, " ")
		report := make([]int, len(nums))
		for j, str := range nums {
			report[j] = aocutils.Stoi(str)
		}

		if isReportSafe(report, -1) {
			p1++
			p2++
		} else if isToleratedReportSafe(report) {
			p2++
		}
	}
	fmt.Println(p1, p2)
}
