package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	l "github.com/sirupsen/logrus"
)

type direction [2]int

func (d direction) dx() int {
	return d[0]
}

func (d direction) dy() int {
	return d[1]
}

var input [][]byte

// true if same data
// false if boundary or other cells
func checkPeri(x int, y int, data string) bool {

	if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) {
		return false
	}

	if string(input[y][x]) == data {
		return true
	} else {
		return false
	}
}

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open(os.Args[1])
	noErr(err)
	defer file.Close()

	input = readInput(file)

	checked := make(map[string]struct{})

	var check func(x, y int, data string, area *int, peri *int, shape *[][2]int)
	check = func(x, y int, data string, area *int, peri *int, shape *[][2]int) {

		if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) {
			return
		}

		if string(input[y][x]) != data {
			return
		}

		if _, ok := checked[fmt.Sprintf("(%d,%d)", x, y)]; !ok {

			if string(input[y][x]) == data {
				*area++
				*shape = append(*shape, [2]int{x, y})

				if !checkPeri(x, y-1, data) {
					*peri++
				}
				if !checkPeri(x+1, y, data) {
					*peri++
				}
				if !checkPeri(x, y+1, data) {
					*peri++
				}
				if !checkPeri(x-1, y, data) {
					*peri++
				}
			}

			checked[fmt.Sprintf("(%d,%d)", x, y)] = struct{}{}
		} else {
			// possible faliure to read some spots?
			// or possible infinite loop if removed?
			// anyway it worked
			return
		}

		check(x, y-1, data, area, peri, shape)
		check(x+1, y, data, area, peri, shape)
		check(x, y+1, data, area, peri, shape)
		check(x-1, y, data, area, peri, shape)

	}

	// number of sides of a shape always equals number of corners in the shape.

	cost := 0
	costWithSides := 0

	for y := range input {
		for x := range input[y] {
			area := 0
			peri := 0

			shape := make([][2]int, 0)

			check(x, y, string(input[y][x]), &area, &peri, &shape)

			if area != 0 && peri != 0 {
				// l.Debugf("Point (%d, %d), data : %s, area: %d, peri: %d", x, y, string(input[y][x]), area, peri)

				// l.Debugf("shape with point (%d, %d) is %v", x, y, shape)
			}
			cost = cost + area*peri
			data := string(input[y][x])

			if x == 0 && y == 0 {
				fmt.Println("0,0")
			}
			sides := getSides(shape, data)

			if area != 0 && sides != 0 {
				l.Debugf("Point (%d, %d), data : %s, area: %d, sides: %d", x, y, string(input[y][x]), area, sides)
			}
			costWithSides += area * sides
		}
	}

	l.Infoln("total cost (area * perimeter): ", cost)
	l.Infoln("total cost (area * sides): ", costWithSides)
}
func getSides(shape [][2]int, data string) int {

	dir := []direction{
		{0, 1},   // 0 or North + clockwise
		{1, 1},   // 1
		{1, 0},   // 2
		{1, -1},  // 3
		{0, -1},  // 4
		{-1, -1}, // 5
		{-1, 0},  // 6
		{-1, 1},  // 7
	}

	edges := [][2]int{
		{0, 1}, //directions
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 0},
	}

	inwCorners := [][2]int{
		{0, 1}, // edges
		{2, 3},
		{4, 5},
		{6, 7},
	}

	// outer corners
	outerCorners := 0
	inwardCorners := 0
	for _, s := range shape {

		x, y := s[0], s[1]

		down := checkPeri(x, y+1, data)
		right := checkPeri(x+1, y, data)
		up := checkPeri(x, y-1, data)
		left := checkPeri(x-1, y, data)

		openSides := 0

		if !up {
			openSides++
		}
		if !right {
			openSides++
		}
		if !down {
			openSides++
		}
		if !left {
			openSides++
		}

		if openSides == 4 {
			//single cell
			outerCorners += 4
		}

		if openSides == 3 {
			outerCorners += 2
		}

		if openSides == 2 {

			if (up == down) || (left == right) {
				outerCorners += 0
			} else {
				outerCorners += 1
			}
		}

		for _, c := range inwCorners {

			edge1, edge2 := edges[c[0]], edges[c[1]]

			dir1, dir2 := dir[edge1[0]], dir[edge1[1]]
			_, dir4 := dir[edge2[0]], dir[edge2[1]]

			if checkPeri(x+dir1.dx(), y+dir1.dy(), data) && !checkPeri(x+dir2.dx(), y+dir2.dy(), data) && checkPeri(x+dir4.dx(), y+dir4.dy(), data) {

				inwardCorners += 1
			}

		}

		l.Debugln(x, y, " inner: ", inwardCorners, "outerCorners: ", outerCorners)
	}

	return outerCorners + inwardCorners
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
