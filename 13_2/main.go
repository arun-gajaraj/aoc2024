package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	file, err := os.Open(os.Args[1])
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

/*
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=10000000008400, Y=10000000005400
*/

// https://www.youtube.com/watch?v=-5J-DAsWuJc
func check(m machine, costs *[]int) {

	ax, ay := m.a.dx, m.a.dy
	bx, by := m.b.dx, m.b.dy
	px, py := m.prize.x, m.prize.y

	if ax*by == ay*bx {
		fmt.Println("ax*by == ay*bx")
		return
	}

	aCount := float64((px*by - py*bx)) / float64((ax*by - ay*bx))
	bCount := float64((float64(px) - float64(ax)*aCount)) / float64(bx)

	fmt.Println(aCount, bCount)

	if aCount != math.Trunc(aCount) || bCount != math.Trunc(bCount) {
		return
	}

	*costs = append(*costs, int(aCount*3+bCount))

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
				x: px + 10000000000000,
				y: py + 10000000000000,
			},
		})

	}

	return input

}
