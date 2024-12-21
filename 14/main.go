package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	l "github.com/sirupsen/logrus"
)

type robot struct {

	//initial positions
	ix int
	iy int

	//velocity
	vx int
	vy int

	//current positio
	x int
	y int
}

const (
	height = 103
	width  = 101

	seconds = 100
)

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open("input.txt")
	noErr(err)
	defer file.Close()

	robots := make([]*robot, 0)
	robots = readInput(file)

	for _, r := range robots {
		nx, ny := (r.ix+r.vx*seconds)%width, (r.iy+r.vy*seconds)%height
		fmt.Println(nx, ny)

		for nx < 0 {
			nx = nx + width
		}
		for ny < 0 {
			ny = ny + height
		}
		nx = nx % width
		ny = ny % height

		r.x, r.y = nx, ny
	}

	first, second, third, fourth := 0, 0, 0, 0

	var display [height][width]int

	for _, r := range robots {
		// fmt.Printf("%+v", *r)

		display[r.y][r.x] += 1

		if r.x < (width/2) && r.y < (height/2) {
			first++
		}
		if r.x > (width/2) && r.y < (height/2) {
			second++
		}
		if r.x < (width/2) && r.y > (height/2) {
			third++
		}
		if r.x > (width/2) && r.y > (height/2) {
			fourth++
		}

	}

	fmt.Println(first, second, third, fourth)

	for c := range display {
		for r := range display[c] {

			if display[c][r] != 0 {

				fmt.Print(display[c][r])
			} else {

				fmt.Print(".")
			}
		}
		fmt.Println("")
	}

	fmt.Println("security factor :", first*second*third*fourth)

}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(r io.Reader) []*robot {

	var input []*robot
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := string(scanner.Bytes())
		var r robot

		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.ix, &r.iy, &r.vx, &r.vy)
		input = append(input, &r)
	}
	return input

}
