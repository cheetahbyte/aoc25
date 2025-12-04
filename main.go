package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	_ "github.com/cheetahbyte/aoc25/day01"
	_ "github.com/cheetahbyte/aoc25/day02"
	_ "github.com/cheetahbyte/aoc25/day03"
	_ "github.com/cheetahbyte/aoc25/day04"
	"github.com/cheetahbyte/aoc25/registry"
)

func measureAndPrint[T any](label string, fn func() T) {
	runtime.GC()

	var mStart, mEnd runtime.MemStats
	runtime.ReadMemStats(&mStart)

	startTime := time.Now()
	res := fn()
	elapsed := time.Since(startTime)

	runtime.ReadMemStats(&mEnd)

	allocDiff := int64(mEnd.Alloc) - int64(mStart.Alloc)
	totalAllocDiff := int64(mEnd.TotalAlloc) - int64(mStart.TotalAlloc)

	fmt.Printf("%s: %-15v\t(%v)  alloc: %dB  totalAlloc: %dB\n",
		label, res, elapsed, allocDiff, totalAllocDiff)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		for _, e := range registry.All() {
			measureAndPrint(e.Label, e.Fn)
		}
		return
	}

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

	found := false
	for _, e := range registry.All() {
		if e.Day == day && (part == 0 || e.Part == part) {
			measureAndPrint(e.Label, e.Fn)
			found = true
		}
	}

	if !found {
		if part > 0 {
			fmt.Printf("No entry found for day %d part %d\n", day, part)
		} else {
			fmt.Printf("No entries found for day %d\n", day)
		}
	}
}
