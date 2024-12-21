package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	wordlen    = 4
	searchText = "XMAS"
)

type grid [][]byte
type word []point

type point struct {
	x int
	y int
}

type direction struct {
	dx int
	dy int
}

func (g grid) rowLen() int {
	return len(g[0])
}

func (g grid) columnLen() int {
	return len(g)
}

func (g grid) pointsOfInterest() (poi []point) {
	for y, row := range g {
		for x, letter := range row {

			if string(letter) == "X" {
				poi = append(poi, point{x, y})
			}
		}
	}
	return
}

func (p *point) directions() []direction {
	return []direction{
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}
}

func (p *point) traceAll(g grid) []word {
	var words []word

	for _, dir := range p.directions() {
		w := p.lookupWord(dir, g)
		words = append(words, w...)
	}
	fmt.Printf("Found All : %v\n", words)
	return words
}

func (p *point) outOfLimits(g grid) bool {
	if p.x < 0 || p.x >= g.rowLen() || p.y < 0 || p.y >= g.columnLen() {
		return true
	}
	return false
}

func (p *point) lookupWord(d direction, g grid) []word {

	var findText string
	var w word
	var words []word

	for i := 0; i < wordlen; i++ {
		spot := point{
			x: p.x + i*d.dx,
			y: p.y + i*d.dy,
		}
		// fmt.Println("checking point: ", spot, "direction", d)

		if spot.outOfLimits(g) {
			return nil
		}

		findText = findText + string(g[spot.y][spot.x])
		w = append(w, spot)
	}

	if findText == searchText {
		fmt.Printf("Found at point: %v, direction: %v, word: %s\n", p, d, findText)
		words = append(words, w)
	} else {
		return nil
	}

	return words
}

func main() {

	f, err := os.Open("input4.txt")
	// f, err := os.Open("input4.txt")
	noErr(err)

	input := readInput(f)

	points := input.pointsOfInterest()

	fmt.Println("Points of interest: ", points)

	count := 0
	for _, poi := range points {
		count = count + len(poi.traceAll(input))
	}

	fmt.Printf("Total Count: %d\n", count)
}

func readInput(reader io.Reader) (g grid) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lineBytes := make([]byte, len(scanner.Bytes()))
		copy(lineBytes, scanner.Bytes())
		g = append(g, lineBytes)
	}
	// fmt.Println(g)
	return
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}
