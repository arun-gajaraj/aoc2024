package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	l "github.com/sirupsen/logrus"
)

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open("input.txt")
	noErr(err)
	defer file.Close()

	grid := readInput(file)

	antennae := make(map[string][][]int)

	//unique locations of antinodes
	uniqueLocs := make(map[string]struct{})

	addNode := func(x, y int) {
		uniqueLocs[fmt.Sprintf("(%d, %d)", x, y)] = struct{}{}
		l.Debugf("added antinode: (%d, %d)", x, y)
	}

	for i, line := range grid {
		for j := range line {
			spot := string(grid[i][j])
			if spot != "." {
				antennae[spot] = append(antennae[spot], []int{j, i})
			}
		}
	}

	l.Debugln(antennae)

	for antType, locations := range antennae {

		for first := 0; first < len(locations); first++ {
			for second := 1; second < len(locations); second++ {

				l.Debugf("antType : %s, Points: (%d,%d) and (%d,%d)", antType, locations[first][0], locations[first][1], locations[second][0], locations[second][1])

				pointA := locations[first]
				pointB := locations[second]

				dx, dy := pointB[0]-pointA[0], pointB[1]-pointA[1]

				if dx == 0 && dy == 0 {
					continue
				}

				addNode(pointA[0], pointA[1])
				addNode(pointB[0], pointB[1])

				newPointX := pointB[0] + dx
				newPointY := pointB[1] + dy
				for newPointX >= 0 && newPointX < len(grid[0]) && newPointY >= 0 && newPointY < len(grid) {

					addNode(newPointX, newPointY)

					newPointX = newPointX + dx
					newPointY = newPointY + dy
				}

				newPointX = pointA[0] - dx
				newPointY = pointA[1] - dy
				for newPointX >= 0 && newPointX < len(grid[0]) && newPointY >= 0 && newPointY < len(grid) {

					addNode(newPointX, newPointY)

					newPointX = newPointX - dx
					newPointY = newPointY - dy
				}
			}
		}
	}

	l.Infof("Length: %d, List := %v\n", len(uniqueLocs), uniqueLocs)

}

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
