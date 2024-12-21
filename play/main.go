package main

import "fmt"

func main() {

	display := [][]int{
		{0, 1, 0, 1, 0},
		{1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0},
	}
	tree := true
	width := 3
	for row := range display {
		for col := range display[row] {

			if col <= width/2 {

				fmt.Println("comparing: ", col, width-col+1)
				if display[row][col] != display[row][(width-col+1)] {
					tree = false
				}
			}
		}

	}
	if tree {
		fmt.Println("still true")
	}
}
