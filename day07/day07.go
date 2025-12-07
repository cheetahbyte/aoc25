package day07

import (
	"fmt"
	"slices"

	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

func IndexAll(input []string) (indices []int) {
	for i, v := range input {
		if v == "^" {
			indices = append(indices, i)
		}
	}
	return
}

func Part1() any {
	data := util.GetGrid()
	d := *data

	splitCount := 0

	starter := d[0]
	starter_pos := slices.Index(starter, "S")
	fmt.Println("Starter Position found at", starter_pos)

	currentPipePositions := []int{starter_pos}

	d = d[1:]

	for i := range d {
		nextPositionsMap := make(map[int]bool)

		for _, pos := range currentPipePositions {
			if pos < 0 || pos >= len(d[i]) {
				continue
			}

			char := d[i][pos]

			if char == "^" {
				splitCount++
				nextPositionsMap[pos-1] = true
				nextPositionsMap[pos+1] = true
			} else {
				// no splitter - extend beam
				nextPositionsMap[pos] = true

				d[i][pos] = "|"
			}
		}

		currentPipePositions = []int{}
		for k := range nextPositionsMap {
			currentPipePositions = append(currentPipePositions, k)
		}
	}

	return splitCount
}

func Part2() any {
	data := util.GetGrid()
	d := *data

	starter := d[0]
	starterPos := slices.Index(starter, "S")

	currentTimelines := make(map[int]uint64)
	currentTimelines[starterPos] = 1

	grid := d[1:]

	for r := range grid {
		nextTimelines := make(map[int]uint64)
		rowWidth := len(grid[r])

		for pos, count := range currentTimelines {
			if pos < 0 || pos >= rowWidth {
				continue
			}

			char := grid[r][pos]

			if char == "^" {
				if pos-1 >= 0 {
					nextTimelines[pos-1] += count
				}
				if pos+1 < rowWidth {
					nextTimelines[pos+1] += count
				}
			} else {
				nextTimelines[pos] += count
			}
		}
		currentTimelines = nextTimelines
	}

	var totalTimelines uint64 = 0
	for _, count := range currentTimelines {
		totalTimelines += count
	}

	return totalTimelines
}

func init() {
	util.ReadFileGrid("day07/input.txt")
	registry.Register(7, 1, "D07P1", Part1)
	registry.Register(7, 2, "D07P2", Part2)
}
