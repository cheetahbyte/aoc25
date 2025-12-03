package day03

import (
	"strconv"
	"strings"

	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

func GetSum(targetLength int, banks *[]string) int64 {
	var sum int64 = 0
	for _, bank := range *banks {
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

func Part1() any {
	data := util.GetData()
	const targetLength = 2
	return GetSum(targetLength, data)
}

func Part2() any {
	data := util.GetData()
	const targetLength = 12
	return GetSum(targetLength, data)
}

func init() {
	util.ReadFile("day03/input.txt")
	registry.Register(2, 1, "D03P1", Part1)
	registry.Register(2, 2, "D03P2", Part2)
}
