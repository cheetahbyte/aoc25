package day02

import (
	"math"
	"strconv"
	"strings"

	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

type Range struct {
	Start int
	End   int
}

var ranges []Range

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

func halvesMatchMath(n int) bool {
	if n < 0 {
		return false
	}

	if n < 10 {
		return false
	}

	digits := int(math.Log10(float64(n))) + 1

	if digits%2 != 0 {
		return false
	}

	halfLen := digits / 2
	divisor := int(math.Pow10(halfLen))

	upperHalf := n / divisor
	lowerHalf := n % divisor

	return upperHalf == lowerHalf
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

func betterRepetition(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	for i := 1; i <= length/2; i++ {
		if length%i == 0 {
			match := true
			for j := i; j < length; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}

	return false
}

func Part1() any {
	sum := 0
	for _, chunk := range ranges {
		for i := chunk.Start; i <= chunk.End; i++ {
			if halvesMatch(i) {
				sum += i
			}
		}
	}
	return sum
}

func Part2() any {
	sum := 0
	for _, chunk := range ranges {
		for i := chunk.Start; i <= chunk.End; i++ {
			if findRepetitions(i) {
				sum += i
			}
		}
	}

	return sum
}

func Part1Math() any {
	sum := 0
	for _, chunk := range ranges {
		for i := chunk.Start; i <= chunk.End; i++ {
			if halvesMatchMath(i) {
				sum += i
			}
		}
	}
	return sum
}

func Part2Rep() any {
	sum := 0
	for _, chunk := range ranges {
		for i := chunk.Start; i <= chunk.End; i++ {
			if betterRepetition(i) {
				sum += i
			}
		}
	}

	return sum
}

func parseRanges() {
	data := *util.GetData()
	chunked := []Range{}
	line := strings.SplitSeq(data[0], ",")
	for l := range line {
		split := strings.Split(l, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		chunked = append(chunked, Range{Start: start, End: end})
	}
	ranges = chunked
}

func init() {
	util.ReadFile("day02/input.txt")
	parseRanges()
	registry.Register(2, 1, "D02P1", Part1)
	registry.Register(2, 2, "D02P2", Part2)
	registry.Register(2, 3, "D02P1+", Part1Math)
	registry.Register(2, 4, "D02P2+", Part2Rep)
}
