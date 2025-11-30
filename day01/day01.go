package day01

/*
#include <stdlib.h>
#include "part2.h"
*/
import "C"
import (
	"unsafe"

	"github.com/cheetahbyte/aoc25/registry"
	"github.com/cheetahbyte/aoc25/util"
)

func parseData() []string {
	lines := util.ReadFile("day01/input.txt")
	return lines
}

func Part1() any {
	return 0
}

func Part2() any {
	data := parseData()
	cStrArray := make([]*C.char, len(data))
	for i, s := range data {
		cStr := C.CString(s)
		cStrArray[i] = cStr
		defer C.free(unsafe.Pointer(cStr))
	}

	ptrToFirstElem := (**C.char)(unsafe.Pointer(&cStrArray[0]))
	count := C.int(len(data))

	return C.Part2_Bridge(ptrToFirstElem, count)
}

func init() {
	registry.Register(1, 1, "D01P1", Part1)
	registry.Register(1, 2, "D01P2", Part2)
}
