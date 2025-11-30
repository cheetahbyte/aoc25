package day01

import (
	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

func parseData() []string {
	lines := util.ReadFile("day01/input.txt")
	return lines
}

func Part1() any {
	return 0
}

func Part2() any {
	return 0
}

func init() {
	registry.Register(1, 1, "D01P1", Part1)
	registry.Register(1, 2, "D01P2", Part2)
}
