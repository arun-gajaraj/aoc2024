package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	l "github.com/sirupsen/logrus"
)

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(r io.Reader) [][]byte {

	var input [][]byte
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lineCopy := make([]byte, len(scanner.Bytes()))
		copy(lineCopy, scanner.Bytes())
		input = append(input, lineCopy)
	}
	return input

}

type spot struct {
	x int
	y int
}
type direction struct {
	dx int
	dy int
}

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open("input.txt")
	noErr(err)
	defer file.Close()

	input := readInput(file)

	potentialObstacles := make([]spot, 0)
	var initX, initY int

	// get all potential spots for placing obstacles
	for line := range input {
		for col := range input[line] {
			if string(input[line][col]) == "." {
				potentialObstacles = append(potentialObstacles, spot{col, line})
			}
			if string(input[line][col]) == "^" {
				initX, initY = col, line
			}
		}
	}

	loopCount := 0
	for _, newObs := range potentialObstacles {
		// newObs := spot{0, 0}
		grid := clone(input)
		grid[newObs.y][newObs.x] = []byte("0")[0]

		// l.Debugln("obstacle in ", newObs.x, newObs.y)

		if good(grid, initX, initY) {
			loopCount++
			// l.Debugf("%v", newObs)
		}
	}

	fmt.Printf("%v", loopCount)
}

func good(grid [][]byte, posX int, posY int) bool {

	been := make(map[string]struct{})
	directions := []direction{
		{0, -1}, // up
		{1, 0},
		{0, 1},
		{-1, 0},
	}
	d := 0

	for {

		// fmt.Println(posX, posY, d)

		newPosX := posX + directions[d].dx
		newPosY := posY + directions[d].dy

		if _, ok := been[toString(posX, posY, d)]; ok {
			// l.Debugln(posX, posY, d)
			return true
		}

		been[toString(posX, posY, d)] = struct{}{}

		if newPosX < 0 || newPosX >= len(grid[0]) || newPosY < 0 || newPosY >= len(grid) {
			// fmt.Println(newPosX, newPosY, d)
			return false
		}

		if string(grid[newPosY][newPosX]) == "#" || string(grid[newPosY][newPosX]) == "0" {
			d = (d + 1) % 4
		} else {
			posX, posY = newPosX, newPosY

		}

	}
}

func clone(inp [][]byte) (cloned [][]byte) {
	for row := range inp {
		lineCopy := make([]byte, len(inp[row]))
		copy(lineCopy, inp[row])
		cloned = append(cloned, lineCopy)
	}
	return
}

func toString(i, j, k int) string {
	return fmt.Sprintf("%d-%d-%d", i, j, k)
}
