package day06

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

func Part1() any {
	var parsed = [][]string{}
	data := util.GetData()

	for _, v := range *data {
		f := []string{}
		for v := range strings.SplitSeq(v, " ") {
			v := strings.Trim(v, " ")
			if len(v) > 0 {
				f = append(f, v)
			}
		}
		parsed = append(parsed, f)
	}

	sum := 0
	if len(parsed) == 0 {
		return 0
	}

	numCols := len(parsed[0])

	for x := 0; x < numCols; x++ {
		numbers := []int{}
		var operator rune

		for y := 0; y < len(parsed); y++ {
			if x >= len(parsed[y]) {
				continue
			}

			val := parsed[y][x]
			num, err := strconv.Atoi(val)

			if err != nil {
				operator = []rune(val)[0]
			} else {
				numbers = append(numbers, num)
			}
		}

		columnResult := 0

		if len(numbers) > 0 {
			columnResult = numbers[0]
			for _, n := range numbers[1:] {
				switch operator {
				case '+':
					columnResult += n
				case '*':
					columnResult *= n
				}
			}
		}

		sum += columnResult
	}

	return sum
}

func Part2() any {
	rawRows := *util.GetData()

	maxWidth := 0
	for _, row := range rawRows {
		if len(row) > maxWidth {
			maxWidth = len(row)
		}
	}

	// Transpose Logic (Rows -> Columns)
	cols := make([][]rune, maxWidth)
	for x := 0; x < maxWidth; x++ {
		col := make([]rune, len(rawRows))
		for y := 0; y < len(rawRows); y++ {
			if x < len(rawRows[y]) {
				col[y] = rune(rawRows[y][x])
			} else {
				col[y] = ' '
			}
		}
		cols[x] = col
	}

	grandTotal := 0

	currentNumbers := []int{}
	var currentOperator rune = 0

	// Iterate Columns Right-to-Left
	for x := len(cols) - 1; x >= 0; x-- {
		col := cols[x]

		digits := []rune{}
		columnIsSpace := true

		for _, char := range col {
			if char != ' ' {
				columnIsSpace = false
				if unicode.IsDigit(char) {
					digits = append(digits, char)
				} else if char == '+' || char == '*' {
					currentOperator = char
				}
			}
		}

		if len(digits) > 0 {
			val, _ := strconv.Atoi(string(digits))
			currentNumbers = append(currentNumbers, val)
		}

		if columnIsSpace || x == 0 {
			if len(currentNumbers) > 0 && currentOperator != 0 {
				blockResult := currentNumbers[0]

				// Apply the operator to the subsequent numbers
				for _, n := range currentNumbers[1:] {
					switch currentOperator {
					case '+':
						blockResult += n
					case '*':
						blockResult *= n
					}
				}
				grandTotal += blockResult
			}

			currentNumbers = []int{}
			currentOperator = 0
		}
	}

	return grandTotal
}

func init() {
	util.ReadFile("day06/input.txt")
	util.ReadFileGrid("day06/input.txt")
	registry.Register(6, 1, "D06P1", Part1)
	registry.Register(6, 2, "D06P2", Part2)
}
