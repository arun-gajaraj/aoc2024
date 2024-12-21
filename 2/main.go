package main

import (
	"aoc/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	safeReports             int
	safeReportsWithDampener int
)

func main() {
	inputFile, err := os.Open("2/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {

		// data collect and initialise
		line := scanner.Text()
		results := strings.Split(line, " ")

		numbers := make([]int, 0, 5)

		for _, v := range results {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			numbers = append(numbers, n)
		}

		fmt.Printf("Reports : %v\n", numbers)

		// numbers from report ready.

		if delta1to3(numbers) {
			if isIncreasing(numbers) || isDecreasing(numbers) {
				safeReports++
			}
		}

		if safeAfterDampening(numbers) {
			fmt.Printf("Number of safe reports without dampener: %d", safeReports)
		}

	}

	fmt.Printf("Number of safe reports WITH dampener: %d\n", safeReportsWithDampener)
}

func delta1to3(list []int) bool {

	// logrus.WithField("report", list).Infoln("Finding delta")

	var prevLevel int
	for i, level := range list {

		// first element
		if i == 0 {
			prevLevel = list[i]
			continue
		}

		// Delta
		delta := helpers.AbsInt(level - prevLevel)
		// fmt.Printf("%d diff %d is %d\n", level, prevLevel, delta)
		if delta < 1 || delta > 3 {
			return false
		}

		// last element
		if i == len(list)-1 {
			break
		}
		prevLevel = list[i]
	}

	// logrus.Infoln("delta true")
	return true
}

func isIncreasing(list []int) bool {
	for i := range list {
		if i == 0 {
			continue
		}

		if list[i-1] < list[i] {
			continue
		} else {
			return false
		}

	}
	return true
}
func isDecreasing(list []int) bool {
	for i := range list {
		if i == 0 {
			continue
		}

		if list[i-1] > list[i] {
			continue
		} else {
			return false
		}

	}
	return true
}

func safeAfterDampening(input []int) bool {

	for i := range input {
		var temp []int

		for j := range input {
			if i == j {
				continue
			}

			temp = append(temp, input[j])
		}

		fmt.Println(temp)
		// Check if safe
		if delta1to3(temp) {
			if isIncreasing(temp) || isDecreasing(temp) {
				return true
			}
		}
	}

	return false
}
