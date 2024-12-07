package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type grid [][]byte

type dir struct {
	x int
	y int
}

const word = "XMAS"

func main() {

	file, err := os.Open("input.txt")
	noErr(err)

	var input grid

	bytes, err := io.ReadAll(file)
	fmt.Println(string(bytes))

	data := strings.Split(string(bytes), "\n")

	fmt.Println(len(data))

	for _, line := range data {

		input = append(input, []byte(line))
	}

	fmt.Println(len(input), len(input[0]))

	directions := []dir{
		{0, -1},  //up
		{1, -1},  // up right
		{1, 0},   // right
		{1, 1},   //down right
		{0, 1},   //down
		{-1, 1},  //down left
		{-1, 0},  // left
		{-1, -1}, //up left
	}
	count := 0

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {

			for _, d := range directions {

				found := true
				for w := 0; w < len(word); w++ {
					rowIndex := row + (d.x * w)
					colIndex := col + (d.y * w)

					if rowIndex < 0 || rowIndex >= len(input) || colIndex < 0 || colIndex >= len(input[row]) {
						found = false
						break
					}

					fmt.Println(row, col, d, string(word[w]), rowIndex, colIndex)
					if string(word[w]) != string(input[colIndex][rowIndex]) {
						found = false
						break
					}
				}
				if found {

					// fmt.Println("FOUND")
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}
