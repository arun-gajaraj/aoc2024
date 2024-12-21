package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	l "github.com/sirupsen/logrus"
)

var grid [][]int
var scoremap map[string]struct{}

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open(os.Args[1])
	noErr(err)
	defer file.Close()

	grid = readInput(file)

	trailHeads := make([][]int, 0)

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 {
				trailHeads = append(trailHeads, []int{x, y})
			}
		}
	}

	sum := 0
	scoremap = make(map[string]struct{})

	for _, head := range trailHeads {
		sum = sum + getSumOfTrails(head)
	}

	l.Infoln("sum of all distinct trails: ", sum)

	l.Infof("sum of all trails is : %d", len(scoremap))
	// l.Infof("%+v", scoremap)
}

func getSumOfTrails(spot []int) int {

	dir := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	var score int //all paths not score
	var navigate func([]int, []int, *int, []int)

	navigate = func(spot []int, d []int, score *int, origin []int) {

		x, y := spot[0], spot[1]
		dx, dy := d[0], d[1]

		if x+dx < 0 || x+dx >= len(grid[0]) || y+dy < 0 || y+dy >= len(grid) {
			return
		}

		if grid[y+dy][x+dx]-grid[y][x] != 1 {
			return
		}
		if grid[y+dy][x+dx] == 9 {
			*score = *score + 1
			scoremap[fmt.Sprintf("%d,%d->%d,%d", origin[0], origin[1], x+dx, y+dy)] = struct{}{}
		}

		navigate([]int{x + dx, y + dy}, dir[0], score, origin)
		navigate([]int{x + dx, y + dy}, dir[1], score, origin)
		navigate([]int{x + dx, y + dy}, dir[2], score, origin)
		navigate([]int{x + dx, y + dy}, dir[3], score, origin)
	}

	navigate(spot, dir[0], &score, spot)
	navigate(spot, dir[1], &score, spot)
	navigate(spot, dir[2], &score, spot)
	navigate(spot, dir[3], &score, spot)

	return score
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(r io.Reader) [][]int {
	var input [][]int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lineCopy := make([]byte, len(scanner.Bytes()))
		copy(lineCopy, scanner.Bytes())

		input = append(input, byteToIntSlice(lineCopy))
	}
	return input

}

func byteToIntSlice(in []byte) (out []int) {
	for _, b := range in {
		bStr := string(b)
		num, err := strconv.Atoi(bStr)
		noErr(err)
		out = append(out, num)
	}
	return
}
