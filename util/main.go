package util

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

var data *[]string
var dataGrid *[][]string

func ReadFile(path string) {
	content, err := os.ReadFile(path)
	ErrPanic(err)
	scanner := bufio.NewScanner(bytes.NewReader(content))
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	data = &lines
}

func ReadFileGrid(path string) {
	content, err := os.ReadFile(path)
	ErrPanic(err)
	scanner := bufio.NewScanner(bytes.NewReader(content))
	var gr [][]string
	for scanner.Scan() {
		line := scanner.Text()
		gr = append(gr, strings.Split(line, ""))
	}
	dataGrid = &gr
}

func GetGrid() *[][]string {
	return dataGrid
}

func GetData() *[]string {
	return data
}
