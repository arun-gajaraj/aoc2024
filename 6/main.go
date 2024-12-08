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

type guard struct {
	posX int
	posY int
	dir  direction

	stepsTaken        int
	distinctPositions int

	plot [][]byte
}

type direction struct {
	dx int
	dy int
}

func (g *guard) turnRight() {
	if g.dir.dx == 0 {
		g.dir.dy = -g.dir.dy
	}
	g.dir.dx, g.dir.dy = g.dir.dy, g.dir.dx
}

func (g *guard) obstacleAhead() (obstrcted, exit bool) {

	if g.posX+g.dir.dx < 0 || g.posX+g.dir.dx >= len(g.plot[0]) || g.posY+g.dir.dy < 0 || g.posY+g.dir.dy >= len(g.plot) {

		// for i := range g.plot {
		// 	for j := range g.plot[i] {
		// 		fmt.Print(string(g.plot[i][j]))
		// 	}
		// 	fmt.Println("")
		// }

		// l.WithField("g", fmt.Sprintf("%+v", *g)).Debugln("Crossing boundary")

		return true, true
	}

	ahead := string(g.plot[g.posY+g.dir.dy][g.posX+g.dir.dx])
	if ahead == "#" || ahead == "0" {
		return true, false
	}
	return false, false
}

func (g *guard) stepForward() {

	g.plot[g.posY][g.posX] = []byte("X")[0]
	g.stepsTaken++

	g.posX = g.posX + g.dir.dx
	g.posY = g.posY + g.dir.dy

	if string(g.plot[g.posY][g.posX]) != "X" {
		g.distinctPositions++
	}

}
func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open("sample.txt")
	noErr(err)
	defer file.Close()

	input := readInput(file)

	// row, col := len(input), len(input[0])

	sx, sy := -1, -1
	directions := []direction{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	for row := range input {
		for col := range input {
			if string(input[row][col]) == "^" {
				sx, sy = col, row
			}
		}
	}

	l.Debugln("init pos", sx, sy)

	g := &guard{
		posX: sx,
		posY: sy,

		dir: directions[0],

		distinctPositions: 1,
		plot:              input,
	}

	for {
		if obs, exit := g.obstacleAhead(); obs {

			if exit {
				l.Infof("%+v", *g)
				break
			}
			g.turnRight()
		}

		g.stepForward()

	}

	for line := range g.plot {
		for col := range g.plot[line] {
			fmt.Print(string(g.plot[line][col]))
		}
		fmt.Println("")
	}

}
