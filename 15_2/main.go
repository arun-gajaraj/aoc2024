package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	l "github.com/sirupsen/logrus"
)

var grid [][]string
var directions []string
var iPos [2]int

func printGrid() {
	for _, row := range grid {
		for _, col := range row {
			if col == "@" {
				fmt.Print("\033[31m@\033[0m")
				continue
			}
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

	pos := iPos
	for _, d := range directions {

		time.Sleep(time.Millisecond * 20)
		clearScreen()
		dxy := getDir(d)

		if strings.ContainsAny(d, "<>") {
			pos = moveHzntl(dxy, pos)
		}
		if strings.ContainsAny(d, "v^") {
			pos = moveVrtcl(dxy, pos)
		}

		printGrid()
	}

	sum := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == "[" {
				sum += (100*row + col)
			}
		}
	}

	l.Infoln("sum of all box coordinates: ", sum)
}

func moveVrtcl(dir [2]int, pos [2]int) [2]int {

	nx, ny := pos[0]+dir[0], pos[1]+dir[1]

	if grid[ny][nx] == "#" {
		return pos
	}
	//free to move one step
	if grid[ny][nx] == "." {
		grid[ny][nx] = "@"
		grid[pos[1]][pos[0]] = "."
		return [2]int{nx, ny}
	}

	// immediate box up or down
	firstRow := make([][]int, 0)

	if grid[ny][nx] == "]" && dir[1] == -1 {
		firstRow = append(firstRow, []int{nx, ny}, []int{nx + dir[1], ny})
	} else if grid[ny][nx] == "]" && dir[1] == 1 {
		firstRow = append(firstRow, []int{nx, ny}, []int{nx - dir[1], ny})
	} else if grid[ny][nx] == "[" && dir[1] == 1 {
		firstRow = append(firstRow, []int{nx, ny}, []int{nx + dir[1], ny})
	} else if grid[ny][nx] == "[" && dir[1] == -1 {
		firstRow = append(firstRow, []int{nx, ny}, []int{nx - dir[1], ny})
	}

	var immovable bool
	train := append([][]int{}, firstRow...)

	for _, pt := range firstRow {
		add(pt, dir, &train, &immovable)
	}
	if immovable {
		return pos
	}

	replace(train, dir)
	grid[ny][nx] = "@"
	grid[pos[1]][pos[0]] = "."

	//TODO
	return [2]int{nx, ny}

}

func add(point []int, dir [2]int, train *[][]int, immovable *bool) {

	dx, dy := dir[0], dir[1]

	nextRow := make([][]int, 0)

	px, py := point[0], point[1]
	nx, ny := px+dx, py+dy

	if grid[ny][nx] == "." {
		return
	}
	if grid[ny][nx] == "#" {
		*immovable = true
		return
	}

	if grid[ny][nx] == grid[py][px] {
		nextRow = append(nextRow, []int{nx, ny})
	}

	if grid[ny][nx] != grid[py][px] {
		nextRow = append(nextRow, []int{nx, ny})

		if grid[ny][nx] == "]" && dy == -1 {
			nextRow = append(nextRow, []int{nx, ny}, []int{nx + dy, ny})
		} else if grid[ny][nx] == "]" && dy == 1 {
			nextRow = append(nextRow, []int{nx, ny}, []int{nx - dy, ny})
		} else if grid[ny][nx] == "[" && dy == 1 {
			nextRow = append(nextRow, []int{nx, ny}, []int{nx + dy, ny})
		} else if grid[ny][nx] == "[" && dy == -1 {
			nextRow = append(nextRow, []int{nx, ny}, []int{nx - dy, ny})
		}

	}

	*train = append(*train, nextRow...)
	for _, nr := range nextRow {
		add(nr, dir, train, immovable)
	}

}

func replace(train [][]int, dir [2]int) {

	temp := [1000][1000]string{}

	for _, t := range train {
		temp[t[1]][t[0]] = grid[t[1]][t[0]]
	}

	// replace with .
	for _, t := range train {
		grid[t[1]][t[0]] = "."
	}

	//replace train with movement
	for _, t := range train {
		grid[t[1]+dir[1]][t[0]] = temp[t[1]][t[0]]
	}

}

func moveHzntl(dir [2]int, pos [2]int) [2]int {

	boxes := ""

	temp := []int{pos[0], pos[1]}
	freeSpace := true

	for {

		tx := temp[0] + dir[0]
		ty := temp[1] + dir[1]

		if grid[ty][tx] == "#" {
			freeSpace = false
			break
		}

		if grid[ty][tx] == "]" || grid[ty][tx] == "[" {

			boxes += grid[ty][tx]

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
	var popped string

	tx, ty := pos[0]+dir[0], pos[1]+dir[1]

	for len(boxes) > 0 {

		boxes, popped = popStr(boxes)
		grid[ty][tx] = popped

		tx, ty = tx+dir[0], ty+dir[1]

	}
	grid[endPos[1]][endPos[0]] = "@"

	return pos

}

func popStr(in string) (out, popped string) {
	if len(in) == 1 {
		return "", in
	}

	return in[1:], string(in[0])
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
			for p, i := range line {

				if strings.Contains(string(i), "#") {
					lineStr = append(lineStr, "#", "#")
				}
				if strings.Contains(string(i), "O") {
					lineStr = append(lineStr, "[", "]")
				}
				if strings.Contains(string(i), ".") {
					lineStr = append(lineStr, ".", ".")
				}
				if strings.Contains(string(i), "@") {

					if string(i) == "@" {
						initPos = [2]int{p * 2, lineNum}
					}
					lineStr = append(lineStr, "@", ".")
				}
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

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
