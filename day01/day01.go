package day01

import (
	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

const (
	dialSize       = 100
	startPos       = 50
	directionRight = 82
)

type instruction struct {
	right bool
	steps int
}

var instructions []instruction

func parseInstructions() []instruction {
	data := util.GetData()
	result := make([]instruction, len(*data))

	for i, v := range *data {
		right := v[0] == directionRight

		steps := 0
		for _, b := range v[1:] {
			if b < '0' || b > '9' {
				break
			}
			steps = steps*10 + int(b-'0')
		}

		result[i] = instruction{right: right, steps: steps}
	}
	return result
}

func Part1() any {
	pos := startPos
	hits := 0
	for _, instruction := range instructions {
		dir, steps := instruction.right, instruction.steps
		if dir {
			pos += steps
		} else {
			pos -= steps
		}
		pos %= dialSize
		if pos < 0 {
			pos += dialSize
		}

		if pos == 0 {
			hits++
		}
	}
	return hits
}

func Part2() any {
	pos := startPos
	hits := 0

	for _, ins := range instructions {
		if ins.right {
			total := pos + ins.steps
			hits += total / dialSize
			pos = total % dialSize
		} else {
			target := pos - ins.steps
			if target <= 0 {
				blockHits := (-target / dialSize) + 1
				if pos == 0 {
					blockHits--
				}
				hits += blockHits
			}

			pos = target % dialSize
			if pos < 0 {
				pos += dialSize
			}
		}
	}

	return hits
}

func init() {
	util.ReadFile("day01/input.txt")
	instructions = parseInstructions()
	registry.Register(1, 1, "D01P1", Part1)
	registry.Register(1, 2, "D01P2", Part2)
}
