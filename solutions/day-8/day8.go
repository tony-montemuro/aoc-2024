package main

import (
	"fmt"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

func pointInBounts(p *aocutils.Point, width, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

func pointToInt(p *aocutils.Point, width int) int {
	return (p.Y * width) + p.X
}

func findAntinodeLocations(antennas map[rune][]aocutils.Point, width, height int, isHarmonics bool) int {
	antinodes := make(map[int]bool)
	for _, points := range antennas {
		for i, p1 := range points {
			for j := i + 1; j < len(points); j++ {
				p2 := points[j]
				dx, dy := p1.X-p2.X, p1.Y-p2.Y
				a1, a2 := aocutils.Point{X: p1.X + dx, Y: p1.Y + dy}, aocutils.Point{X: p2.X - dx, Y: p2.Y - dy}
				if isHarmonics {
					antinodes[pointToInt(&p1, width)] = true
					antinodes[pointToInt(&p2, width)] = true

					for pointInBounts(&a1, width, height) || pointInBounts(&a2, width, height) {
						if pointInBounts(&a1, width, height) {
							antinodes[pointToInt(&a1, width)] = true
							a1.X += dx
							a1.Y += dy
						}
						if pointInBounts(&a2, width, height) {
							antinodes[pointToInt(&a2, width)] = true
							a2.X -= dx
							a2.Y -= dy
						}
					}
				} else {
					if pointInBounts(&a1, width, height) {
						antinodes[pointToInt(&a1, width)] = true
					}
					if pointInBounts(&a2, width, height) {
						antinodes[pointToInt(&a2, width)] = true
					}
				}
			}
		}
	}
	return len(antinodes)
}

func main() {
	input := aocutils.GetInputRows()
	width, height := len(input[0]), len(input)
	antennas := make(map[rune][]aocutils.Point)
	for y, row := range input {
		for x, t := range row {
			if t != '.' {
				antennas[t] = append(antennas[t], aocutils.Point{X: x, Y: y})
			}
		}
	}

	p1, p2 := findAntinodeLocations(antennas, width, height, false), findAntinodeLocations(antennas, width, height, true)
	fmt.Println(p1, p2)
}
