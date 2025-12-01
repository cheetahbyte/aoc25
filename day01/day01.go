package day01

import (
	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func Part1() any {
	data := util.GetData()
	var currentPosition = 50
	var zeroCounter = 0
	for _, v := range *data {
		direction := v[0]
		var steps = 0
		for _, b := range v[1:] {
			if b >= '0' && b <= '9' {
				steps = steps*10 + int(b-'0')
			}
		}
		var newPosition int
		// R is 82 in binary
		if direction == 82 {
			newPosition = mod((currentPosition + steps), 100)
		} else {
			newPosition = mod((currentPosition - steps), 100)
		}
		if newPosition == 0 {
			zeroCounter += 1
		}
		// fmt.Println("took", steps, "to", direction, "started", currentPosition, "ended", newPosition)
		currentPosition = newPosition
	}
	return zeroCounter
}

func Part2() any {
	data := util.GetData()
	var currentPosition = 50
	var zeroCounter = 0

	for _, v := range *data {
		direction := v[0]
		var steps = 0
		for _, b := range v[1:] {
			if b >= '0' && b <= '9' {
				steps = steps*10 + int(b-'0')
			}
		}

		// R is 82 in binary
		if direction == 82 {
			hits := (currentPosition + steps) / 100
			zeroCounter += hits
			currentPosition = (currentPosition + steps) % 100
		} else {
			target := currentPosition - steps
			if target <= 0 {
				hits := (-target / 100) + 1
				if currentPosition == 0 {
					hits -= 1
				}
				zeroCounter += hits
			}
			currentPosition = ((currentPosition-steps)%100 + 100) % 100
		}
	}
	return zeroCounter
}

func init() {
	util.ReadFile("day01/input.txt")
	registry.Register(1, 1, "D01P1", Part1)
	registry.Register(1, 2, "D01P2", Part2)
}
