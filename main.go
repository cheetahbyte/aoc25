package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	_ "github.com/cheetahbyte/aoc25/day01"
	_ "github.com/cheetahbyte/aoc25/day02"
	_ "github.com/cheetahbyte/aoc25/day03"
	_ "github.com/cheetahbyte/aoc25/day04"
	_ "github.com/cheetahbyte/aoc25/day05"
	_ "github.com/cheetahbyte/aoc25/day06"
	"github.com/cheetahbyte/aoc25/registry"
)

func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%d ns", d.Nanoseconds())
	} else if d < time.Millisecond {
		return fmt.Sprintf("%.2f µs", float64(d.Nanoseconds())/1000)
	} else if d < time.Second {
		return fmt.Sprintf("%.2f ms", d.Seconds()*1000)
	}
	return fmt.Sprintf("%.2f s", d.Seconds())
}

func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

func measureAndPrint[T any](label string, fn func() T, repeat int) {
	var (
		totalElapsed     time.Duration
		minElapsed       time.Duration = 1<<63 - 1 // Max possible duration
		maxElapsed       time.Duration = 0
		totalBytes       uint64
		totalAllocations uint64
		lastRes          T
	)

	// Warmup
	fn()

	for i := 0; i < repeat; i++ {
		runtime.GC()

		var mStart, mEnd runtime.MemStats
		runtime.ReadMemStats(&mStart)

		startTime := time.Now()
		lastRes = fn()
		elapsed := time.Since(startTime)

		runtime.ReadMemStats(&mEnd)

		// Update Min/Max
		if elapsed < minElapsed {
			minElapsed = elapsed
		}
		if elapsed > maxElapsed {
			maxElapsed = elapsed
		}

		totalElapsed += elapsed
		totalBytes += (mEnd.TotalAlloc - mStart.TotalAlloc)
		totalAllocations += (mEnd.Mallocs - mStart.Mallocs)
	}

	avgElapsed := totalElapsed / time.Duration(repeat)
	avgBytes := totalBytes / uint64(repeat)
	avgAllocs := totalAllocations / uint64(repeat)
	fmt.Printf("%-10s  %10s (%s-%s)  %10s (%d allocs)\n",
		label,
		formatDuration(avgElapsed),
		formatDuration(minElapsed),
		formatDuration(maxElapsed),
		formatBytes(avgBytes),
		avgAllocs)

	fmt.Printf("  -> %v\n\n", lastRes)
}

type filter struct {
	// map[Day]map[Part]bool
	dayParts map[int]map[int]bool
}

func (f filter) matches(e registry.Entry) bool {
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

func parseFilter(inputs []string) (filter, error) {
	f := filter{dayParts: map[int]map[int]bool{}}

	for _, raw := range inputs {
		parts := strings.SplitSeq(raw, ",")
		for p := range parts {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}

			segments := strings.Split(p, ".")
			if len(segments) > 2 {
				return filter{}, fmt.Errorf("invalid filter format: %q", p)
			}

			day, err := strconv.Atoi(segments[0])
			if err != nil {
				return filter{}, fmt.Errorf("invalid day: %q", segments[0])
			}

			if _, ok := f.dayParts[day]; !ok {
				f.dayParts[day] = map[int]bool{}
			}

			if len(segments) == 2 {
				part, err := strconv.Atoi(segments[1])
				if err != nil {
					return filter{}, fmt.Errorf("invalid part: %q", segments[1])
				}
				f.dayParts[day][part] = true
			}
		}
	}
	return f, nil
}

// --- Main ---

func main() {
	list := flag.Bool("list", false, "list available solutions")
	runFlag := flag.String("run", "", "comma-separated list of day or day.part (e.g. 1,2.1)")
	repeat := flag.Int("repeat", 1, "number of times to run for averaging")
	flag.Parse()

	if *repeat < 1 {
		fmt.Println("Error: repeat must be at least 1")
		os.Exit(1)
	}

	entries := registry.All()

	// Sort entries so Day 1 Part 1 always appears before Day 1 Part 2, etc.
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Day != entries[j].Day {
			return entries[i].Day < entries[j].Day
		}
		return entries[i].Part < entries[j].Part
	})

	if *list {
		fmt.Println("Available Solutions:")
		for _, e := range entries {
			fmt.Printf("  Day %02d Part %d: %s\n", e.Day, e.Part, e.Label)
		}
		return
	}

	// Combine -run flag and positional args into a single filter list
	// This allows: `go run . -run 1` OR `go run . 1`
	filterInputs := flag.Args()
	if *runFlag != "" {
		filterInputs = append(filterInputs, *runFlag)
	}

	runFilter, err := parseFilter(filterInputs)
	if err != nil {
		fmt.Printf("Error parsing filter: %v\n", err)
		os.Exit(1)
	}

	found := false
	for _, e := range entries {
		if runFilter.matches(e) {
			measureAndPrint(e.Label, e.Fn, *repeat)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching solutions found.")
	}
}
