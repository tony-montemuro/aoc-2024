package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func findStartPosition(grid *aocutils.Grid) (aocutils.Point, error) {
	gridData := grid.Data

	for y := 1; y < len(gridData)-2; y++ {
		for x := 1; x < len(gridData[y])-2; x++ {
			char := gridData[y][x]
			if char == '^' {
				grid.UpdateDataAt(x, y, '.')
				return aocutils.Point{X: x, Y: y}, nil
			}
		}
	}

	return aocutils.Point{X: 0, Y: 0}, fmt.Errorf("start position not found")
}

func traverse(grid *aocutils.Grid, obstaclePos int) ([]int, error) {
	seenPoints := make(map[int][]aocutils.Direction)
	exited := false

	for !exited {
		position := grid.GetIntegerPoint()
		prevDirections, exists := seenPoints[position]
		if exists && slices.Contains(prevDirections, grid.Direction) {
			return nil, fmt.Errorf("loop detected")
		}
		seenPoints[position] = append(seenPoints[position], grid.Direction)

		nextChar := grid.LookUp()
		for nextChar == "#" || grid.GetUpPosition() == obstaclePos {
			grid.RotateRight(1)
			nextChar = grid.LookUp()
		}

		err := grid.MoveUp()
		if err != nil {
			exited = true
		}
	}

	points := []int{}
	for p := range seenPoints {
		points = append(points, p)
	}
	return points, nil
}

func countLoops(grid *aocutils.Grid, visitedPoints []int, startingPosition aocutils.Point) int {
	loopsCtr := 0
	for _, point := range visitedPoints {
		grid.Position = startingPosition
		grid.Direction = aocutils.North

		_, err := traverse(grid, point)
		if err != nil {
			loopsCtr++
		}
	}
	return loopsCtr
}

func main() {
	input := aocutils.GetInputRows()

	grid := aocutils.NewGrid(1, 1, 0, input)
	startingPos, err := findStartPosition(grid)
	if err != nil {
		log.Fatal(err.Error())
	}
	grid.SetPosition(startingPos)

	visitedPoints, _ := traverse(grid, -1)
	ctr := countLoops(grid, visitedPoints, startingPos)
	fmt.Println(len(visitedPoints), ctr)
}
