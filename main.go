package main

import (
	"fmt"
	"runtime"
	"time"

	_ "github.com/cheetahbyte/aoc25/day01"
	_ "github.com/cheetahbyte/aoc25/day02"
	_ "github.com/cheetahbyte/aoc25/day03"
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

	// Calculate differences
	allocDiff := int64(mEnd.Alloc) - int64(mStart.Alloc)
	totalAllocDiff := int64(mEnd.TotalAlloc) - int64(mStart.TotalAlloc)

	fmt.Printf("%s: %-15v\t(%v)  alloc: %dB  totalAlloc: %dB\n",
		label, res, elapsed, allocDiff, totalAllocDiff)
}

func main() {
	for _, e := range registry.All() {
		measureAndPrint(e.Label, e.Fn)
	}
}
