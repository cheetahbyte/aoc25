package day04

import (
	"strings"

	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

var dirs = [][2]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func parseGrid(data []string) ([][]string, int, int) {
	grid := [][]string{}
	for _, row := range data {
		if row == "" {
			continue
		}
		split := strings.Split(row, "")
		grid = append(grid, split)
	}
	return grid, len(grid), len(grid[0])
}

func countNeighbors(grid [][]string, x, y, rows, cols int) int {
	count := 0
	for _, d := range dirs {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && nx < cols && ny >= 0 && ny < rows {
			if grid[ny][nx] == "@" {
				count++
			}
		}
	}
	return count
}

func Part1() any {
	data := *util.GetData()
	grid, rows, cols := parseGrid(data)
	accessibleCount := 0
	for y := range rows {
		for x := range cols {
			if grid[y][x] != "@" {
				continue
			}

			paperNeighbors := countNeighbors(grid, x, y, rows, cols)

			if paperNeighbors < 4 {
				accessibleCount++
			}
		}
	}

	return accessibleCount
}

func Part2() any {
	data := *util.GetData()
	grid, rows, cols := parseGrid(data)
	totalRemoved := 0

	for {
		toRemove := [][2]int{}

		for y := range rows {
			for x := range cols {
				if grid[y][x] == "@" {
					if countNeighbors(grid, x, y, rows, cols) < 4 {
						toRemove = append(toRemove, [2]int{x, y})
					}
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}
		// remove them
		totalRemoved += len(toRemove)
		for _, coord := range toRemove {
			grid[coord[1]][coord[0]] = "."
		}
	}

	return totalRemoved
}

func init() {
	util.ReadFile("day04/input.txt")
	registry.Register(4, 1, "D04P1", Part1)
	registry.Register(4, 2, "D04P2", Part2)
}
