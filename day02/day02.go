package day02

import (
	"strconv"
	"strings"

	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

type Range struct {
	Start int
	End   int
}

func halvesMatch(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	if length%2 != 0 {
		return false
	}

	mid := length / 2

	part1 := s[:mid]
	part2 := s[mid:]

	return part1 == part2
}

func findRepetitions(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	for i := 1; i <= length/2; i++ {
		if length%i == 0 {
			part := s[0:i]

			reps := length / i

			expectation := strings.Repeat(part, reps)

			if expectation == s {
				return true
			}
		}
	}

	return false
}

func Part1() any {
	data := *util.GetData()
	chunked := []Range{}
	line := strings.Split(data[0], ",")
	sum := 0
	for _, l := range line {
		split := strings.Split(l, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		chunked = append(chunked, Range{Start: start, End: end})
	}
	for _, chunk := range chunked {
		for i := chunk.Start; i <= chunk.End; i++ {
			if halvesMatch(i) {
				sum += i
			}
		}
	}
	return sum
}

func Part2() any {
	data := *util.GetData()
	chunked := []Range{}
	line := strings.Split(data[0], ",")
	sum := 0
	for _, l := range line {
		split := strings.Split(l, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		chunked = append(chunked, Range{Start: start, End: end})
	}
	for _, chunk := range chunked {
		for i := chunk.Start; i <= chunk.End; i++ {
			if findRepetitions(i) {
				sum += i
			}
		}
	}

	return sum
}

func init() {
	util.ReadFile("day02/input.txt")
	registry.Register(2, 1, "D02P1", Part1)
	registry.Register(2, 2, "D02P2", Part2)
}
