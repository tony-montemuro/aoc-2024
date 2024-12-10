package main

import (
	"fmt"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func traverseForPeaks(grid *aocutils.Grid, hits map[int]bool) {
	val := grid.GetValue()
	digit := aocutils.Stoi(val)
	pos := grid.Position

	if val == "9" {
		hits[grid.GetIntegerPoint()] = true
		return
	}

	grid.Position = pos
	if grid.LookUp() != "." && aocutils.Stoi(grid.LookUp()) == digit+1 {
		grid.MoveUp()
		traverseForPeaks(grid, hits)
	}
	grid.Position = pos
	if grid.LookRight() != "." && aocutils.Stoi(grid.LookRight()) == digit+1 {
		grid.MoveRight()
		traverseForPeaks(grid, hits)
	}
	grid.Position = pos
	if grid.LookDown() != "." && aocutils.Stoi(grid.LookDown()) == digit+1 {
		grid.MoveDown()
		traverseForPeaks(grid, hits)
	}
	grid.Position = pos
	if grid.LookLeft() != "." && aocutils.Stoi(grid.LookLeft()) == digit+1 {
		grid.MoveLeft()
		traverseForPeaks(grid, hits)
	}
}

func traverseForPaths(grid *aocutils.Grid, path []int, hits map[string]bool) {
	val := grid.GetValue()
	digit := aocutils.Stoi(val)
	pos := grid.Position
	path = append(path, grid.GetIntegerPoint())

	if val == "9" {
		strPath := ""
		for _, n := range path {
			strPath += string(n)
		}
		hits[strPath] = true
		return
	}

	grid.Position = pos
	if grid.LookUp() != "." && aocutils.Stoi(grid.LookUp()) == digit+1 {
		grid.MoveUp()
		traverseForPaths(grid, path, hits)
	}
	grid.Position = pos
	if grid.LookRight() != "." && aocutils.Stoi(grid.LookRight()) == digit+1 {
		grid.MoveRight()
		traverseForPaths(grid, path, hits)
	}
	grid.Position = pos
	if grid.LookDown() != "." && aocutils.Stoi(grid.LookDown()) == digit+1 {
		grid.MoveDown()
		traverseForPaths(grid, path, hits)
	}
	grid.Position = pos
	if grid.LookLeft() != "." && aocutils.Stoi(grid.LookLeft()) == digit+1 {
		grid.MoveLeft()
		traverseForPaths(grid, path, hits)
	}
}

func countReachablePeaks(p aocutils.Point, grid *aocutils.Grid) int {
	grid.Position = p
	hits := make(map[int]bool)
	traverseForPeaks(grid, hits)
	return len(hits)
}

func countDistinctTrails(p aocutils.Point, grid *aocutils.Grid) int {
	grid.Position = p
	hits := make(map[string]bool)
	path := []int{}
	traverseForPaths(grid, path, hits)
	return len(hits)
}

func main() {
	input := aocutils.GetInputRows()
	grid := aocutils.NewGrid(1, 1, aocutils.North, input)
	p1, p2 := 0, 0

	for y, row := range grid.Data {
		for x, space := range row {
			score := string(space)
			if score == "0" {
				point := aocutils.Point{X: x, Y: y}
				p1 += countReachablePeaks(point, grid)
				p2 += countDistinctTrails(point, grid)
			}
		}
	}
	fmt.Println(p1, p2)
}
