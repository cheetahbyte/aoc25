package main

import (
	"flag"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/cheetahbyte/aoc25/day01"
	_ "github.com/cheetahbyte/aoc25/day02"
	_ "github.com/cheetahbyte/aoc25/day03"
	_ "github.com/cheetahbyte/aoc25/day04"
	"github.com/cheetahbyte/aoc25/registry"
)

func measureAndPrint[T any](label string, fn func() T, repeat int) {
	var (
		totalElapsed    time.Duration
		totalAllocDiff  int64
		totalTotalAlloc int64
		lastRes         T
	)

	for i := 0; i < repeat; i++ {
		runtime.GC()

		var mStart, mEnd runtime.MemStats
		runtime.ReadMemStats(&mStart)

		startTime := time.Now()
		lastRes = fn()
		elapsed := time.Since(startTime)

		runtime.ReadMemStats(&mEnd)

		allocDiff := int64(mEnd.Alloc) - int64(mStart.Alloc)
		totalAlloc := int64(mEnd.TotalAlloc) - int64(mStart.TotalAlloc)

		totalElapsed += elapsed
		totalAllocDiff += allocDiff
		totalTotalAlloc += totalAlloc
	}

	avgElapsed := totalElapsed / time.Duration(repeat)
	avgAlloc := totalAllocDiff / int64(repeat)
	avgTotalAlloc := totalTotalAlloc / int64(repeat)

	plural := "s"
	if repeat == 1 {
		plural = ""
	}

	fmt.Printf("%s: %-15v\t(avg over %d run%s) %v  alloc: %dB  totalAlloc: %dB\n",
		label, lastRes, repeat, plural, avgElapsed, avgAlloc, avgTotalAlloc)
}

type filter struct {
	dayParts map[int]map[int]bool
}

func parseRunFilter(raw string) (filter, error) {
	f := filter{dayParts: map[int]map[int]bool{}}
	if raw == "" {
		return f, nil
	}

	parts := strings.Split(raw, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		segments := strings.Split(p, ".")
		if len(segments) > 2 {
			return filter{}, fmt.Errorf("invalid filter %q", p)
		}

		day, err := strconv.Atoi(segments[0])
		if err != nil {
			return filter{}, fmt.Errorf("invalid day in filter %q", p)
		}

		if _, ok := f.dayParts[day]; !ok {
			f.dayParts[day] = map[int]bool{}
		}

		if len(segments) == 2 {
			part, err := strconv.Atoi(segments[1])
			if err != nil {
				return filter{}, fmt.Errorf("invalid part in filter %q", p)
			}
			f.dayParts[day][part] = true
		}
	}

	return f, nil
}

func matchesFilter(e registry.Entry, f filter) bool {
	if len(f.dayParts) == 0 {
		return true
	}

	parts, ok := f.dayParts[e.Day]
	if !ok {
		return false
	}

	if len(parts) == 0 {
		return true
	}

	return parts[e.Part]
}

func main() {
	list := flag.Bool("list", false, "list available solutions")
	run := flag.String("run", "", "comma-separated list of day or day.part values to run (e.g. 1,2.1,3.2)")
	repeat := flag.Int("repeat", 1, "number of times to run each solution for averaging")

	flag.Parse()

	entries := registry.All()

	if *list {
		for _, e := range entries {
			fmt.Printf("day %02d part %d\t%s\n", e.Day, e.Part, e.Label)
		}
		return
	}

	if *repeat < 1 {
		fmt.Println("repeat must be at least 1")
		return
	}

	f, err := parseRunFilter(*run)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := flag.Args()

	if *run == "" && len(args) > 0 {
		day, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid day")
			return
		}

		var part int
		if len(args) > 1 {
			part, err = strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid part")
				return
			}
		}

		f.dayParts[day] = map[int]bool{}
		if part > 0 {
			f.dayParts[day][part] = true
		}
	}

	found := false
	for _, e := range entries {
		if matchesFilter(e, f) {
			measureAndPrint(e.Label, e.Fn, *repeat)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching entries found")
	}
}
