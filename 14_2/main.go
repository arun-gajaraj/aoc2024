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

	seconds = 1
)

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open("input.txt")
	noErr(err)
	defer file.Close()

	robots := make([]*robot, 0)
	robots = readInput(file)

	t := 0
	of, err := os.Create("tree.txt")
	noErr(err)
	defer of.Close()

	for {

		for _, r := range robots {
			nx, ny := (r.x + r.vx*seconds), (r.y + r.vy*seconds)

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

		var display [height][width]int

		for _, r := range robots {

			display[r.y][r.x] += 1

		}

		for c := range display {
			for r := range display[c] {

				if display[c][r] != 0 {

					fmt.Print(display[c][r])
					fmt.Fprint(of, "#")
				} else {

					fmt.Print(".")
					fmt.Fprint(of, " ")
				}
			}
			fmt.Println("")
			fmt.Fprintln(of, "")
		}

		t++
		fmt.Println(t)
		fmt.Fprintln(of, "@@@@@@@@@@@@@@@@@@@@@@", t)

		if t > 10000 {
			break
		}

	}

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

		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.x, &r.y, &r.vx, &r.vy)
		input = append(input, &r)
	}
	return input

}
