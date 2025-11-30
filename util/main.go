package util

import (
	"bufio"
	"bytes"
	"os"
)

func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(path string) []string {
	content, err := os.ReadFile(path)
	ErrPanic(err)
	scanner := bufio.NewScanner(bytes.NewReader(content))
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
