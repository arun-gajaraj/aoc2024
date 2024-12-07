package main

import (
	"bufio"
	"os"
	"strings"

	l "github.com/sirupsen/logrus"
)

type grid [][]byte

type dir struct {
	dx int
	dy int
}

type point struct {
	x int
	y int
}

func main() {
	l.SetLevel(l.InfoLevel)
	l.SetLevel(l.DebugLevel)
	var input grid

	file, err := os.Open("input4.txt")
	noErr(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCopy := make([]byte, len(scanner.Bytes()))
		copy(lineCopy, scanner.Bytes())
		input = append(input, lineCopy)
	}

	_ = []dir{
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}

	poi := input.pointsOfInterest()

	l.Infoln("Points of Interest", poi)
	l.Info("Count: ", len(poi))

}

func (g *grid) pointsOfInterest() []point {

	pointAs := make([]point, 0)
	for row := range *g {
		for col := range (*g)[row] {
			l.Traceln(g.elemAt(row, col))
			if string(g.elemAt(row, col)) == "A" {

				pointAs = append(pointAs, point{row, col})
			}
		}
	}

	xmasPoints := make([]point, 0)

	for _, p := range pointAs {

		// remove first/last rows/columns since they cut-off at the edge
		if p.x == 0 || p.x == g.rowLen()-1 || p.y == 0 || p.y == g.colLen()-1 {
			continue
		}

		if (g.elemAt(p.x-1, p.y-1) == "M" && g.elemAt(p.x+1, p.y+1) == "S" ||
			g.elemAt(p.x-1, p.y-1) == "S" && g.elemAt(p.x+1, p.y+1) == "M") &&
			(g.elemAt(p.x-1, p.y+1) == "M" && g.elemAt(p.x+1, p.y-1) == "S" ||
				g.elemAt(p.x-1, p.y+1) == "S" && g.elemAt(p.x+1, p.y-1) == "M") {
			xmasPoints = append(xmasPoints, p)
		}
	}

	return xmasPoints

}
func (g *grid) elemAt(x, y int) string {
	return string((*g)[y][x])
}

func (g *grid) isOffLimits(x, y int) bool {
	return x < 0 || x >= len(*g) || y < 0 || y >= len((*g)[0])
}

func (g *grid) rowLen() int {
	return len(*g)
}

func (g *grid) colLen() int {
	return len((*g)[0])
}

func (g *grid) String() string {
	s := new(strings.Builder)
	for _, line := range *g {
		for _, letter := range line {
			s.WriteString(string(letter))
		}
		s.WriteString("\n")
	}
	return s.String()
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}
