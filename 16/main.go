package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	l "github.com/sirupsen/logrus"
)

var grid [][]string

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open(os.Args[1])
	noErr(err)
	defer file.Close()

	var start, end [2]int
	grid, start, end = readInput(file)

	fmt.Println(start, end)

}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(r io.Reader) (grid [][]string, startPos, endPos [2]int) {

	var input [][]string
	scanner := bufio.NewScanner(r)
	row := 0
	for scanner.Scan() {

		line := scanner.Bytes()
		lineStr := make([]string, len(line))

		for i := range line {

			lineStr[i] = string(line[i])
			if lineStr[i] == "S" {
				startPos = [2]int{i, row}
			}
			if lineStr[i] == "E" {
				endPos = [2]int{i, row}
			}
		}

		input = append(input, lineStr)
		row++
	}
	return input, startPos, endPos

}
