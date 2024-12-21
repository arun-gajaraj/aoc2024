package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	l "github.com/sirupsen/logrus"
)

type machine struct {
	a     button
	b     button
	prize struct {
		x int
		y int
	}
}
type button struct {
	dx int
	dy int
}

var machines []machine

const nMax = 100

var logfile io.Writer

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open("input.txt")
	noErr(err)
	defer file.Close()

	machines = readInput(file)
	noErr(err)

	totalCost := 0

	for _, m := range machines {

		costs := make([]int, 0)

		check(m, &costs)

		mCostLow := 0

		for _, c := range costs {
			if mCostLow == 0 {
				mCostLow = c
				continue
			}
			if c < mCostLow {
				mCostLow = c
				continue
			}
		}

		l.Debugf("lowest cost for machine %+v is %d", m, mCostLow)
		totalCost += mCostLow
	}

	l.Infoln("Total Cost :", totalCost)

}

func check(m machine, costs *[]int) {

	for a := 0; a <= nMax; a++ {
		for b := 0; b <= nMax; b++ {
			if a*m.a.dx+b*m.b.dx == m.prize.x && a*m.a.dy+b*m.b.dy == m.prize.y {
				*costs = append(*costs, a*3+b*1)
			}
		}
	}
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(r io.Reader) []machine {

	var input []machine
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {

		line := string(scanner.Bytes())

		ax, ay, bx, by, px, py := 0, 0, 0, 0, 0, 0
		var err error
		if strings.HasPrefix(line, "Button A:") {
			pre1 := strings.TrimPrefix(line, "Button A: X+")
			numstr := strings.Split(strings.TrimSpace(pre1), ", Y+")

			ax, err = strconv.Atoi(numstr[0])
			noErr(err)
			ay, err = strconv.Atoi(numstr[1])
			noErr(err)

		}
		scanner.Scan()
		line = string(scanner.Bytes())
		if strings.HasPrefix(line, "Button B") {
			pre1 := strings.TrimPrefix(line, "Button B: X+")
			numstr := strings.Split(strings.TrimSpace(pre1), ", Y+")

			bx, err = strconv.Atoi(numstr[0])
			noErr(err)
			by, err = strconv.Atoi(numstr[1])
			noErr(err)
		}
		scanner.Scan()
		line = string(scanner.Bytes())
		if strings.HasPrefix(line, "Prize: X") {
			pre1 := strings.TrimPrefix(line, "Prize: X=")
			numstr := strings.Split(strings.TrimSpace(pre1), ", Y=")

			px, err = strconv.Atoi(numstr[0])
			noErr(err)
			py, err = strconv.Atoi(numstr[1])
			noErr(err)
		}
		scanner.Scan()

		input = append(input, machine{
			a: button{
				dx: ax,
				dy: ay,
			},
			b: button{
				dx: bx,
				dy: by,
			},
			prize: struct {
				x int
				y int
			}{
				x: px,
				y: py,
			},
		})

	}

	return input

}
