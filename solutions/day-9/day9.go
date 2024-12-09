package main

import (
	"fmt"
	"slices"

	"github.com/tony-montemuro/aoc-2024/aocutils"
)

type Block struct {
	id     int
	data   []int
	remain int
}

func logDisk(disk []*Block) {
	for _, block := range disk {
		fmt.Println("id:", block.id, "data:", block.data, "remain:", block.remain)
	}
}

func main() {
	input := aocutils.GetRawInput()
	disk := []*Block{}

	for i, id := 0, 0; i < len(input); i += 2 {
		data := slices.Repeat([]int{id}, aocutils.Stoi(string(input[i])))
		remain := 0

		if i < len(input)-1 {
			remain = aocutils.Stoi(string(input[i+1]))
		}

		disk = append(disk, &Block{id: id, data: data, remain: remain})
		id++
	}

	earlistGap := 0
	for disk[earlistGap].id < disk[len(disk)-1].id {
		lastBlock := disk[len(disk)-1]
		data := lastBlock.data

		for len(data) > 0 && disk[earlistGap].id < lastBlock.id {
			block := disk[earlistGap]
			n := min(block.remain, len(data))

			for i := 0; i < n; i++ {
				block.data = append(block.data, data[len(data)-1])
				lastBlock.data = data[:len(data)-1]
				lastBlock.remain++
				block.remain--
				data = lastBlock.data
			}

			if block.remain == 0 {
				earlistGap++
			}

			if len(lastBlock.data) == 0 {
				disk = disk[:len(disk)-1]
			}
		}
	}

	checksum, pos := 0, 0
	for _, block := range disk {
		for _, id := range block.data {
			checksum += id * pos
			pos++
		}
	}

	fmt.Println(checksum)
}
