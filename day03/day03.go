package day03

import (
	"strconv"
	"strings"

	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

func Part1() any {
	data := util.GetData()
	sum := 0

	for _, bank := range *data {
		maxJoltage := -1
		for i := 0; i < len(bank)-1; i++ {
			for j := i + 1; j < len(bank); j++ {
				d1 := int(bank[i] - '0')
				d2 := int(bank[j] - '0')
				currentJoltage := (d1 * 10) + d2
				if currentJoltage > maxJoltage {
					maxJoltage = currentJoltage
				}
			}
		}
		sum += maxJoltage
	}
	return sum
}

func Part2() any {
	data := util.GetData()
	var sum int64 = 0
	const targetLength = 12

	for _, bank := range *data {
		if len(bank) < targetLength {
			continue
		}

		var sb strings.Builder
		currentPos := 0
		needed := targetLength

		for needed > 0 {
			searchEnd := len(bank) - needed

			bestDigit := -1
			bestIndex := -1

			for i := currentPos; i <= searchEnd; i++ {
				digit := int(bank[i] - '0')

				if digit > bestDigit {
					bestDigit = digit
					bestIndex = i
				}
			}

			sb.WriteString(strconv.Itoa(bestDigit))
			currentPos = bestIndex + 1
			needed--
		}

		val, _ := strconv.ParseInt(sb.String(), 10, 64)
		sum += val
	}
	return sum
}

func init() {
	util.ReadFile("day03/input.txt")
	registry.Register(2, 1, "D03P1", Part1)
	registry.Register(2, 2, "D03P2", Part2)
}
