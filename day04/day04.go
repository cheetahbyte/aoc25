package day04

import (
	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

var neighborOffsets []int

func setupOffsets(width int) {
	neighborOffsets = []int{
		-width - 1, -width, -width + 1,
		-1, 1,
		width - 1, width, width + 1,
	}
}

func parseGrid(raw [][]string) ([]byte, int, int, []int) {
	rows := len(raw)
	cols := len(raw[0])
	grid := make([]byte, rows*cols)
	var active []int

	for y := range rows {
		for x := range cols {
			if raw[y][x] == "@" {
				idx := y*cols + x
				grid[idx] = 1
				active = append(active, idx)
			}
		}
	}
	return grid, cols, rows, active
}

func Part1() any {
	rawGrid := *util.GetGrid()
	grid, width, height, activeIndices := parseGrid(rawGrid)
	setupOffsets(width)

	totalSize := width * height
	accessibleCount := 0

	for _, idx := range activeIndices {
		neighbors := 0
		for _, offset := range neighborOffsets {
			nIdx := idx + offset
			if nIdx >= 0 && nIdx < totalSize {
				cx, cy := idx%width, idx/width
				nx, ny := nIdx%width, nIdx/width
				if (nx-cx) > 1 || (cx-nx) > 1 || (ny-cy) > 1 || (cy-ny) > 1 {
					continue
				}

				if grid[nIdx] == 1 {
					neighbors++
				}
			}
		}

		if neighbors < 4 {
			accessibleCount++
		}
	}

	return accessibleCount
}

func Part2() any {
	raw := *util.GetGrid()
	rows, cols := len(raw), len(raw[0])

	// pad the grid, so every cell has 8 neighbors
	w := cols + 1
	grid := make([]byte, (rows+2)*w)

	// q holds indicies to check
	q := make([]int, 0, rows*cols)

	for y := range rows {
		// off by 1 row to center
		offset := (y + 1) * w
		for x := range cols {
			if raw[y][x] == "@" {
				idx := offset + x
				grid[idx] = 3 // Binary 11: Alive(1) | Queued(2)
				q = append(q, idx)
			}
		}
	}

	offs := []int{-w - 1, -w, -w + 1, -1, 1, w - 1, w, w + 1}

	totalRemoved := 0
	deaths := make([]int, 0, 64)

	// simulation loops
	for len(q) > 0 {
		deaths = deaths[:0]
		for _, idx := range q {
			grid[idx] &= 1
			n := 0
			for _, d := range offs {
				n += int(grid[idx+d] & 1)
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
			if grid[idx]&1 == 1 {
				grid[idx] = 0 // Kill cell
				totalRemoved++

				for _, d := range offs {
					nIdx := idx + d
					if grid[nIdx] == 1 {
						grid[nIdx] = 3
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
