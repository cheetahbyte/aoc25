package day01

/*
#include <stdlib.h>
#include "part2.h"
*/
import "C"
import (
	"strconv"
	"strings"

	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

func parseData() []string {
	lines := util.ReadFile("day01/input.txt")
	return lines
}

type Operation struct {
	Direction string
	Steps     int
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func Part1() any {
	data := parseData()
	var currentPosition = 50
	var zeroCounter = 0
	for _, v := range data {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		direction := string(v[0])
		steps, _ := strconv.Atoi(v[1:])
		var newPosition int
		if direction == "R" {
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
	data := parseData()
	var currentPosition = 50
	var zeroCounter = 0

	for _, v := range data {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		direction := string(v[0])
		steps, _ := strconv.Atoi(v[1:])

		if direction == "R" {
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
	registry.Register(1, 1, "D01P1", Part1)
	registry.Register(1, 2, "D01P2", Part2)
}
