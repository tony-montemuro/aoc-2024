package aocutils

import (
	"fmt"
	"slices"
	"strings"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Point struct {
	X, Y int
}

type Grid struct {
	Data                []string
	Position            Point
	Direction           Direction
	Width               int
	Height              int
	horizontalOobLength int
	verticalOobLength   int
}

func NewGrid(horizontal, vertical int, direction Direction, grid []string) *Grid {
	width, height := len(grid[0]), len(grid)
	paddedGrid := []string{}

	for _, row := range grid {
		pad := strings.Repeat(".", horizontal)
		paddedGrid = append(paddedGrid, pad+row+pad)
	}

	paddedRow := strings.Repeat(".", (2*horizontal)+len(grid[0]))
	for i := 0; i < vertical; i++ {
		paddedGrid = slices.Insert(paddedGrid, 0, paddedRow)
		paddedGrid = append(paddedGrid, paddedRow)
	}

	return &Grid{
		Data:                paddedGrid,
		Direction:           direction,
		Position:            Point{horizontal, vertical},
		Width:               width,
		Height:              height,
		horizontalOobLength: horizontal,
		verticalOobLength:   vertical,
	}
}

func (g *Grid) getMovedCoordinates(direction Direction) (int, int) {
	var dx int
	var dy int
	trueDirection := (g.Direction + direction) % (West + 1)

	switch trueDirection {
	case North:
		dx, dy = 0, -1
	case East:
		dx, dy = 1, 0
	case South:
		dx, dy = 0, 1
	case West:
		dx, dy = -1, 0
	}

	return g.Position.X + dx, g.Position.Y + dy
}

func (g *Grid) move(direction Direction) error {
	x, y := g.getMovedCoordinates(direction)
	if !g.isInGrid(x, y) {
		return fmt.Errorf("this motion leaves the grid")
	}
	g.Position = Point{x, y}
	return nil
}

func (g *Grid) isInGrid(x, y int) bool {
	return x >= g.horizontalOobLength &&
		x < len(g.Data[x])-g.horizontalOobLength &&
		y >= g.verticalOobLength &&
		y < len(g.Data)-g.verticalOobLength
}

func (g *Grid) pointToInt(p Point) int {
	return (g.Width * p.Y) + p.X
}

func (g *Grid) SetPosition(p Point) {
	g.Position = p
}

func (g *Grid) UpdateDataAt(x, y int, newByte byte) {
	g.Data[y] = UpdateCharAt(g.Data[y], rune(newByte), x)
}

func (g *Grid) GetValue() string {
	return string(g.Data[g.Position.Y][g.Position.X])
}

func (g *Grid) LookUp() string {
	movedX, movedY := g.getMovedCoordinates(North)
	return string(g.Data[movedY][movedX])
}

func (g *Grid) LookRight() string {
	movedX, movedY := g.getMovedCoordinates(East)
	return string(g.Data[movedY][movedX])
}

func (g *Grid) LookDown() string {
	movedX, movedY := g.getMovedCoordinates(South)
	return string(g.Data[movedY][movedX])
}

func (g *Grid) LookLeft() string {
	movedX, movedY := g.getMovedCoordinates(West)
	return string(g.Data[movedY][movedX])
}

func (g *Grid) GetUpPosition() int {
	movedX, movedY := g.getMovedCoordinates(North)
	return g.pointToInt(Point{movedX, movedY})
}

func (g *Grid) GetRightPosition() int {
	movedX, movedY := g.getMovedCoordinates(East)
	return g.pointToInt(Point{movedX, movedY})
}

func (g *Grid) GetDownPosition() int {
	movedX, movedY := g.getMovedCoordinates(South)
	return g.pointToInt(Point{movedX, movedY})
}

func (g *Grid) GetLeftPosition() int {
	movedX, movedY := g.getMovedCoordinates(West)
	return g.pointToInt(Point{movedX, movedY})
}

func (g *Grid) MoveUp() error {
	err := g.move(North)
	return err
}

func (g *Grid) MoveRight() error {
	err := g.move(East)
	return err
}

func (g *Grid) MoveDown() error {
	err := g.move(South)
	return err
}

func (g *Grid) MoveLeft() error {
	err := g.move(West)
	return err
}

func (g *Grid) RotateRight(n int) {
	directionCount := int(West) + 1
	g.Direction = Direction((int(g.Direction) + n) % directionCount)
}

func (g *Grid) RotateLeft(n int) {
	directionCount := int(West) + 1
	transformedDirection := (int(g.Direction) - n) % directionCount
	if transformedDirection < 0 {
		transformedDirection += n
	}
	g.Direction = Direction(transformedDirection)
}

func (g *Grid) GetIntegerPoint() int {
	return g.pointToInt(g.Position)
}
