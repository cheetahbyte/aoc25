package day05

import (
	"sort"
	"strconv"
	"strings"

	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

type Interval struct {
	Start int
	End   int
}

func SplitOnce(input []string) ([]string, []string) {
	for i, str := range input {
		if str == "" {
			return input[:i], input[i+1:]
		}
	}
	return input, []string{}
}

func Part1() any {
	data := *util.GetData()
	freshLines, ingrLines := SplitOnce(data)

	ranges := make([]Interval, 0, len(freshLines))

	for _, v := range freshLines {
		before, after, found := strings.Cut(v, "-")
		if !found {
			continue
		}

		start, _ := strconv.Atoi(before)
		end, _ := strconv.Atoi(after)

		ranges = append(ranges, Interval{Start: start, End: end})
	}

	sumFresh := 0

	for _, valStr := range ingrLines {
		val, _ := strconv.Atoi(valStr)

		for _, r := range ranges {
			if val >= r.Start && val <= r.End {
				sumFresh++
				break
			}
		}
	}

	return sumFresh
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var result []Interval
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		if current.End >= next.Start { // Overlap
			if next.End > current.End {
				current.End = next.End
			}
		} else {
			result = append(result, current)
			current = next
		}
	}
	result = append(result, current)
	return result
}

func Part2() any {
	data := *util.GetData()
	freshLines, _ := SplitOnce(data)

	var ranges []Interval
	for _, v := range freshLines {
		before, after, found := strings.Cut(v, "-")
		if !found {
			continue
		}
		start, _ := strconv.Atoi(before)
		end, _ := strconv.Atoi(after)
		ranges = append(ranges, Interval{Start: start, End: end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	if len(ranges) == 0 {
		return 0
	}

	var merged []Interval
	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]

		if next.Start <= current.End {
			if next.End > current.End {
				current.End = next.End
			}
		} else {
			merged = append(merged, current)
			current = next
		}
	}
	merged = append(merged, current)

	totalFresh := 0
	for _, r := range merged {
		totalFresh += (r.End - r.Start) + 1
	}

	return totalFresh
}

func init() {
	util.ReadFile("day05/input.txt")
	registry.Register(5, 1, "D05P1", Part1)
	registry.Register(5, 2, "D05P2", Part2)
}
