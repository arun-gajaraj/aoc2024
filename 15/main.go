package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	l "github.com/sirupsen/logrus"
)

var grid [][]string
var directions []string
var iPos [2]int

func printGrid() {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open("input.txt")
	noErr(err)
	defer file.Close()

	grid, directions, iPos = readInput(file)

	printGrid()

	pos := iPos
	for _, d := range directions {

		dxy := getDir(d)

		pos = move(dxy, pos)

	}

	sum := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == "O" {
				sum += (100*row + col)
			}
		}
	}

	printGrid()
	l.Infoln("sum of all box coordinates: ", sum)
}

func move(dir [2]int, pos [2]int) [2]int {

	boxes := 0

	temp := []int{pos[0], pos[1]}
	freeSpace := true

	for {

		tx := temp[0] + dir[0]
		ty := temp[1] + dir[1]

		if grid[ty][tx] == "#" {
			freeSpace = false
			break
		}

		if grid[ty][tx] == "O" {
			boxes++

			temp[0], temp[1] = tx, ty
			continue
		}

		if grid[ty][tx] == "." {

			temp[0], temp[1] = tx, ty
			break
		}

	}

	if !freeSpace {
		return pos
	}

	// current pos as free space
	grid[pos[1]][pos[0]] = "."

	//temp as end of moving line
	pos = [2]int{pos[0] + dir[0], pos[1] + dir[1]}
	endPos := pos

	tx, ty := pos[0]+dir[0], pos[1]+dir[1]

	for boxes > 0 {

		grid[ty][tx] = "O"
		boxes--

		tx, ty = tx+dir[0], ty+dir[1]

	}
	grid[endPos[1]][endPos[0]] = "@"

	return pos

}

func getDir(symbol string) [2]int {

	if symbol == "^" {
		fmt.Println("up")
		return [2]int{0, -1}
	}
	if symbol == "v" {
		fmt.Println("down")
		return [2]int{0, 1}
	}

	if symbol == "<" {
		fmt.Println("left")
		return [2]int{-1, 0}
	}

	if symbol == ">" {
		fmt.Println("right")
		return [2]int{1, 0}
	}

	return [2]int{}
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(r io.Reader) ([][]string, []string, [2]int) {

	var input [][]string
	var directions []string
	var initPos [2]int

	scanner := bufio.NewScanner(r)
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Bytes()

		if strings.Contains(string(line), "#") {

			var lineStr []string
			for i := range line {

				if string(line[i]) == "@" {

					initPos = [2]int{i, lineNum}
				}
				lineStr = append(lineStr, string(line[i]))

			}
			input = append(input, lineStr)
		}

		if string(line) == "\n" {
			continue
		}

		if strings.ContainsAny(string(line), "^>v<") {

			for _, d := range line {

				directions = append(directions, string(d))
			}
		}
		lineNum++
	}
	return input, directions, initPos

}
