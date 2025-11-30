package main

import (
	"fmt"
	"time"

	"github.com/cheetahbyte/aoc25/day01"
)

func measureAndPrint[T any](label string, fn func() T) {
	startTime := time.Now()
	res := fn()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("%s: %-15v\t(%v)\n", label, res, elapsed)
}

func main() {
	measureAndPrint("D01P1", day01.Part1)
}
