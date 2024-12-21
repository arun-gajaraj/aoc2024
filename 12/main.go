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

	file, err := os.Open(os.Args[1])
	noErr(err)
	defer file.Close()

	input := readInput(file)

	checked := make(map[string]struct{})

	checkPeri := func(x int, y int, data string) bool {

		if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) {
			return false
		}

		if string(input[y][x]) == data {
			return true
		} else {
			return false
		}
	}

	var check func(x, y int, data string, area *int, peri *int)
	check = func(x, y int, data string, area *int, peri *int) {

		if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) {
			return
		}

		if string(input[y][x]) != data {
			return
		}

		if _, ok := checked[fmt.Sprintf("(%d,%d)", x, y)]; !ok {

			if string(input[y][x]) == data {
				*area++

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

		check(x, y-1, data, area, peri)
		check(x+1, y, data, area, peri)
		check(x, y+1, data, area, peri)
		check(x-1, y, data, area, peri)

	}

	cost := 0

	for y := range input {
		for x := range input[y] {
			area := 0
			peri := 0

			check(x, y, string(input[y][x]), &area, &peri)
			if area != 0 && peri != 0 {
				l.Debugf("Point (%d, %d), data : %s, area: %d, peri: %d", x, y, string(input[y][x]), area, peri)
			}
			cost = cost + area*peri
		}
	}

	l.Infoln("total cost: ", cost)
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
