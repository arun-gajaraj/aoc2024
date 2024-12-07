package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	logfile *os.File
)

func main() {

	file, err := os.Open("./1/input.txt")
	if err != nil {
		panic(err)
	}

	logfile, err = os.Create("logfile.log")
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	scanner := bufio.NewScanner(file)

	ListA := make([]int, 0, 100)
	ListB := make([]int, 0, 100)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		a, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		ListA = append(ListA, a)

		b, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		ListB = append(ListB, b)
	}

	sort.IntSlice(ListA).Sort()
	sort.IntSlice(ListB).Sort()

	totalDistance := 0

	for i := range ListA {
		a := ListA[i]
		b := ListB[i]

		temp := a - b

		if temp >= 0 {
			totalDistance = temp + totalDistance
			continue
		} else {
			temp = -temp
			totalDistance = temp + totalDistance
			continue
		}
	}

	fmt.Printf("Total Distances is %d\n", totalDistance)

	fmt.Println("Finding Similarity score now..")

	score := 0
	for _, value := range ListA {
		score = score + value*GetNumberOfOccurances(&ListB, value)
	}

	fmt.Printf("Similarity score is %d\n", score)

}

func GetNumberOfOccurances(list *[]int, elem int) int {

	times := 0

	for _, value := range *list {
		if value == elem {
			times++
		}
	}
	return times
}
