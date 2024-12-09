package main

import (
	"fmt"
	"slices"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

type Block struct {
	id   int
	data []int
	size int
}

func logDisk(disk []*Block) {
	for _, block := range disk {
		fmt.Println("id:", block.id, "data:", block.data, "size:", block.size)
	}
}

func checksum(disk []*Block) int {
	checksum, pos := 0, 0

	for _, block := range disk {
		for _, id := range block.data {
			checksum += id * pos
			pos++
		}
	}

	return checksum
}

func fragment(disk []*Block) int {
	earlistGap := 0
	for disk[earlistGap].id < disk[len(disk)-1].id {
		lastBlock := disk[len(disk)-1]
		data := lastBlock.data

		for len(data) > 0 && disk[earlistGap].id < lastBlock.id {
			block := disk[earlistGap]
			n := min(block.size-len(block.data), len(data))

			for i := 0; i < n; i++ {
				block.data = append(block.data, data[len(data)-1])
				lastBlock.data = data[:len(data)-1]
				data = lastBlock.data
			}

			if block.size-len(block.data) == 0 {
				earlistGap++
			}

			if len(lastBlock.data) == 0 {
				disk = disk[:len(disk)-1]
			}
		}
	}

	return checksum(disk)
}

func defragment(disk []*Block) int {
	index := len(disk) - 1

	return checksum(disk)
}

func main() {
	input := aocutils.GetRawInput()
	d1 := []*Block{}

	for i, id := 0, 0; i < len(input); i += 2 {
		data := slices.Repeat([]int{id}, aocutils.Stoi(string(input[i])))
		size := len(data)

		if i < len(input)-1 {
			size += aocutils.Stoi(string(input[i+1]))
		}

		d1 = append(d1, &Block{id, data, size})
		id++
	}

	d2 := make([]*Block, len(d1))
	copy(d2, d1)
	p1, p2 := fragment(d1), defragment(d2)
	fmt.Println(p1, p2)
}
