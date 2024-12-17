package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

type Robot struct {
	point    *aocutils.Point
	velocity *aocutils.Point
}

func stringToPoint(str string) *aocutils.Point {
	s := strings.Split(strings.Split(str, "=")[1], ",")
	return &aocutils.Point{X: aocutils.Stoi(s[0]), Y: aocutils.Stoi(s[1])}
}

func main() {
	input := aocutils.GetInputRows()
	robots := []Robot{}
	width, height, p1, isTreeFound := 101, 103, 100, false

	pointToInt := func(p aocutils.Point) int {
		return (width * p.Y) + p.X
	}

	getLargestSequence := func(unique map[int]bool) int {
		points := []int{}
		for point := range unique {
			points = append(points, point)
		}
		sort.Ints(points)

		prev, maxSequence, sequence := math.MinInt, 0, 0
		for _, point := range points {
			if point == prev+1 {
				sequence++
			} else {
				sequence = 0
			}

			if sequence > maxSequence {
				maxSequence = sequence
			}
			prev = point
		}

		return maxSequence
	}

	getSafetyScore := func(robots []Robot) int {
		q1, q2, q3, q4 := 0, 0, 0, 0
		unique := make(map[int]bool)

		for _, robot := range robots {
			if robot.point.X < width/2 {
				if robot.point.Y < height/2 {
					q1++
				}
				if robot.point.Y >= height/2+1 {
					q3++
				}
			}
			if robot.point.X >= width/2+1 {
				if robot.point.Y < height/2 {
					q2++
				}
				if robot.point.Y >= height/2+1 {
					q4++
				}
			}
			unique[pointToInt(*robot.point)] = true
		}

		// if we discover a sequence >= 7, we have found the tree
		maxSequence := getLargestSequence(unique)
		if maxSequence >= 7 {
			isTreeFound = true
		}

		return q1 * q2 * q3 * q4
	}

	for _, row := range input {
		s := strings.Split(row, " ")
		robots = append(robots, Robot{point: stringToPoint(s[0]), velocity: stringToPoint(s[1])})
	}

	for i := 1; !isTreeFound || i <= 100; i++ {
		for _, robot := range robots {
			robot.point.X = aocutils.Modulo(robot.point.X+robot.velocity.X, width)
			robot.point.Y = aocutils.Modulo(robot.point.Y+robot.velocity.Y, height)
		}

		safetyScore := getSafetyScore(robots)
		if i == p1 {
			fmt.Println(safetyScore)
		}
		if isTreeFound {
			fmt.Println(i)
		}
	}
}
