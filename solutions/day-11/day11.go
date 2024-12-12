package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func countDigits(number int) int {
	n, digits := number, 0
	for n > 0 {
		n /= 10
		digits++
	}
	return digits
}

func count(n, depth int, depthMap []map[int]int) int {
	if depth == 0 {
		return 1
	}

	result, exists := depthMap[depth][n]
	if exists {
		return result
	}

	size := countDigits(n)
	if n == 0 {
		result = count(1, depth-1, depthMap)
	} else if size%2 == 0 {
		divisor := int(math.Pow(10, float64(size/2)))
		result += count(n/divisor, depth-1, depthMap)
		result += count(n%divisor, depth-1, depthMap)
	} else {
		result = count(n*2024, depth-1, depthMap)
	}

	depthMap[depth][n] = result
	return result
}

func blink(arrangement []int, depth int, depthMap []map[int]int) int {
	sum := 0
	for _, n := range arrangement {
		sum += count(n, depth, depthMap)
	}
	return sum
}

func main() {
	input := aocutils.GetRawInput()
	arrangement := []int{}
	for _, str := range strings.Split(input, " ") {
		arrangement = append(arrangement, aocutils.Stoi(str))
	}

	d1, d2 := 25, 75
	depthMap := make([]map[int]int, d2+1)
	for i := 0; i <= d2; i++ {
		depthMap[i] = make(map[int]int)
	}
	fmt.Println(blink(arrangement, d1, depthMap), blink(arrangement, d2, depthMap))
}
