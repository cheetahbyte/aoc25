package day04

import (
	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

const (
	cellAlive  = 1
	cellQueued = 2
)

var neighborDeltas = [8][2]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func parseGrid(raw [][]string) (grid []byte, width, height int, active []int) {
	height = len(raw)
	if height == 0 {
		return nil, 0, 0, nil
	}

	width = len(raw[0])
	grid = make([]byte, width*height)
	active = make([]int, 0, width*height/4)

	for y := 0; y < height; y++ {
		row := raw[y]
		for x := 0; x < width; x++ {
			if row[x] == "@" {
				idx := y*width + x
				grid[idx] = cellAlive
				active = append(active, idx)
			}
		}
	}
	return grid, width, height, active
}

func Part1() any {
	rawGrid := *util.GetGrid()
	grid, width, height, active := parseGrid(rawGrid)

	accessibleCount := 0

	for _, idx := range active {
		cx := idx % width
		cy := idx / width

		neighbors := 0
		for _, d := range neighborDeltas {
			nx := cx + d[0]
			ny := cy + d[1]
			if nx < 0 || nx >= width || ny < 0 || ny >= height {
				continue
			}
			nIdx := ny*width + nx
			if grid[nIdx] == cellAlive {
				neighbors++
			}
		}

		if neighbors < 4 {
			accessibleCount++
		}
	}

	return accessibleCount
}

func buildPaddedGrid(width, height int, active []int) (grid []byte, stride int, queue []int) {
	stride = width + 2
	h := height + 2

	grid = make([]byte, stride*h)
	queue = make([]int, 0, len(active))

	for _, idx := range active {
		y := idx / width
		x := idx % width
		pIdx := (y+1)*stride + (x + 1)
		grid[pIdx] = cellAlive | cellQueued
		queue = append(queue, pIdx)
	}

	return grid, stride, queue
}

func neighborOffsetsForStride(stride int) []int {
	return []int{
		-stride - 1, -stride, -stride + 1,
		-1, 1,
		stride - 1, stride, stride + 1,
	}
}

func Part2() any {
	raw := *util.GetGrid()
	_, w, h, active := parseGrid(raw)

	grid, stride, q := buildPaddedGrid(w, h, active)
	offs := neighborOffsetsForStride(stride)

	totalRemoved := 0
	deaths := make([]int, 0, 64)

	for len(q) > 0 {
		deaths = deaths[:0]

		for _, idx := range q {
			grid[idx] &= cellAlive
			if grid[idx]&cellAlive == 0 {
				continue
			}

			n := 0
			for _, d := range offs {
				if grid[idx+d]&cellAlive == cellAlive {
					n++
				}
			}

			if n < 4 {
				deaths = append(deaths, idx)
			}
		}

		if len(deaths) == 0 {
			break
		}

		q = q[:0]

		for _, idx := range deaths {
			if grid[idx]&cellAlive == cellAlive {
				grid[idx] = 0
				totalRemoved++

				for _, d := range offs {
					nIdx := idx + d
					if grid[nIdx] == cellAlive {
						grid[nIdx] = cellAlive | cellQueued
						q = append(q, nIdx)
					}
				}
			}
		}
	}

	return totalRemoved
}

func init() {
	util.ReadFileGrid("day04/input.txt")
	registry.Register(4, 1, "D04P1", Part1)
	registry.Register(4, 2, "D04P2", Part2)
}
