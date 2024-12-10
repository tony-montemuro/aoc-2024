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

func (b *Block) getRemainingSpace() int {
	return b.size - len(b.data)
}

func readDisc(input string) []*Block {
	disk := []*Block{}

	for i, id := 0, 0; i < len(input); i += 2 {
		data := slices.Repeat([]int{id}, aocutils.Stoi(string(input[i])))
		size := len(data)

		if i < len(input)-1 {
			size += aocutils.Stoi(string(input[i+1]))
		}

		disk = append(disk, &Block{id, data, size})
		id++
	}

	return disk
}

func checksum(disk []*Block) int {
	checksum, pos := 0, 0

	for _, block := range disk {
		for _, id := range block.data {
			if id >= 0 {
				checksum += id * pos
			}
			pos++
		}
		pos += block.getRemainingSpace()
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
			n := min(block.getRemainingSpace(), len(data))

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
	for i := len(disk) - 1; i > 0; i-- {
		block := disk[i]
		moved := false
		data := []int{}
		if len(block.data) > 0 {
			diff := false
			data = append(data, block.data[0])
			for j := 1; j < len(block.data) && !diff; j++ {
				if block.data[j] == block.data[j-1] {
					data = append(data, block.data[j])
				} else {
					diff = true
				}
			}
		}
		hasOtherData := len(block.data) > len(data)

		for j := 0; j < i && !moved; j++ {
			currentBlock := disk[j]

			if len(data) <= currentBlock.getRemainingSpace() {
				for _, id := range data {
					currentBlock.data = append(currentBlock.data, id)
					data = data[:len(data)-1]
				}

				if hasOtherData {
					for j := 0; j < len(data); j++ {
						block.data[j] = -1
					}
				} else {
					block.data = data
				}
			}
		}
	}

	return checksum(disk)
}

func main() {
	input := aocutils.GetRawInput()
	d1, d2 := readDisc(input), readDisc(input)
	fmt.Println(fragment(d1), defragment(d2))
}
