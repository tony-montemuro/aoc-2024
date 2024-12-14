package main

import (
	"fmt"
	"sort"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func main() {
	grid := aocutils.NewGrid(1, 1, aocutils.North, aocutils.GetInputRows())

	hasVisited := func(m map[int]bool, key int) bool {
		_, exists := m[key]
		return exists
	}

	getRegionMeasurements := func(grid *aocutils.Grid, region string) (map[int]bool, [][]int) {
		regionCoords := make(map[int]bool)
		edgeCoords := make([][]int, 4)
		for i := 0; i < int(aocutils.West); i++ {
			edgeCoords[i] = []int{}
		}

		var traverseRegion func(grid aocutils.Grid)
		traverseRegion = func(grid aocutils.Grid) {
			regionCoords[grid.GetIntegerPoint()] = true

			// north
			if grid.LookUp() != region {
				edgeCoords[int(aocutils.North)] = append(edgeCoords[int(aocutils.North)], grid.GetIntegerPoint())
			} else if !hasVisited(regionCoords, grid.GetUpPosition()) {
				grid.MoveUp()
				traverseRegion(grid)
				grid.MoveDown()
			}

			// east
			if grid.LookRight() != region {
				edgeCoords[int(aocutils.East)] = append(edgeCoords[int(aocutils.East)], grid.GetIntegerPoint())
			} else if !hasVisited(regionCoords, grid.GetRightPosition()) {
				grid.MoveRight()
				traverseRegion(grid)
				grid.MoveLeft()
			}

			// south
			if grid.LookDown() != region {
				edgeCoords[int(aocutils.South)] = append(edgeCoords[int(aocutils.South)], grid.GetIntegerPoint())
			} else if !hasVisited(regionCoords, grid.GetDownPosition()) {
				grid.MoveDown()
				traverseRegion(grid)
				grid.MoveUp()
			}

			// west
			if grid.LookLeft() != region {
				edgeCoords[int(aocutils.West)] = append(edgeCoords[int(aocutils.West)], grid.GetIntegerPoint())
			} else if !hasVisited(regionCoords, grid.GetLeftPosition()) {
				grid.MoveLeft()
				traverseRegion(grid)
				grid.MoveRight()
			}
		}

		traverseRegion(*grid)
		return regionCoords, edgeCoords
	}

	getPerimeter := func(edgeCoords [][]int) int {
		sides := 0
		for _, points := range edgeCoords {
			sides += len(points)
		}
		return sides
	}

	getNumberOfSides := func(edgeCoords [][]int) int {
		intToPoint := func(i int) aocutils.Point {
			return aocutils.Point{X: i % grid.Width, Y: i / grid.Width}
		}

		countHorizontalSides := func(points []int) int {
			sort.Slice(points, func(i, j int) bool {
				return points[i]-points[j] < 0
			})

			sides, prev := 0, -1
			for _, point := range points {
				if point != prev+1 {
					sides++
				}
				prev = point
			}
			return sides
		}

		countVerticalSides := func(p []int) int {
			points := []aocutils.Point{}
			for _, point := range p {
				points = append(points, intToPoint(point))
			}
			sort.SliceStable(points, func(i, j int) bool {
				return points[i].Y-points[j].Y < 0
			})
			sort.SliceStable(points, func(i, j int) bool {
				return points[i].X-points[j].X < 0
			})

			sides := 0
			prev := aocutils.Point{X: -1, Y: -1}
			for _, point := range points {
				if point.X != prev.X || point.Y != prev.Y+1 {
					sides++
				}
				prev = point
			}

			return sides
		}

		sides := 0
		sides += countHorizontalSides(edgeCoords[int(aocutils.North)])
		sides += countHorizontalSides(edgeCoords[int(aocutils.South)])
		sides += countVerticalSides(edgeCoords[int(aocutils.East)])
		sides += countVerticalSides(edgeCoords[int(aocutils.West)])

		return sides
	}

	getPricing := func(grid *aocutils.Grid, perimeterCalcuation func([][]int) int) int {
		visited := make(map[int]bool)
		price := 0

		for y := grid.StartY; y < grid.Height+grid.StartY; y++ {
			for x := grid.StartX; x < grid.Width+grid.StartX; x++ {
				grid.Position = aocutils.Point{X: x, Y: y}
				if !hasVisited(visited, grid.GetIntegerPoint()) {
					regionCoords, edgeCoords := getRegionMeasurements(grid, grid.GetValue())
					area := len(regionCoords)
					for coord := range regionCoords {
						visited[coord] = true
					}
					perim := perimeterCalcuation(edgeCoords)
					price += area * perim
				}
			}
		}

		return price
	}

	fmt.Println(getPricing(grid, getPerimeter), getPricing(grid, getNumberOfSides))
}
