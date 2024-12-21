package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	l "github.com/sirupsen/logrus"
)

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open("input.txt")
	noErr(err)
	defer file.Close()

	input := readInput(file)
	part1sum := 0

	for _, line := range input {
		result, numbers := scanLine(line)

		// l.Debugln(result, numbers)

		if correct(result, numbers) {
			// l.Debugln("p#1 satifies: ", result, numbers)
			part1sum = part1sum + result
		}

	}

	l.Infoln("Part #1: sum of all correct results: ", part1sum)
	part2sum := 0

	for _, line := range input {
		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		result, numbers := scanLine(lineCopy)

		// l.Debugln(result, numbers)

		good := false
		correctP2(result, numbers, numbers[0], 1, &good, "")

		if good {
			l.Debugln("P#2 satifies: ", result, numbers)
			part2sum = part2sum + result
		}

	}

	l.Infoln("Part #2: sum of all correct results: ", part2sum)

	l.Infoln("sum of both results: ", part1sum+part2sum)

}

// 156: 1 5 6

func correctP2(final int, parts []int, temp int, index int, good *bool, printStr string) {
	if index == len(parts) {
		// fmt.Println(temp, final)
		if temp == final {
			*good = true
			fmt.Println(final, parts, *good, printStr)
		}
		return
	}

	correctP2(final, parts, temp+parts[index], index+1, good, printStr+" +")
	correctP2(final, parts, temp*parts[index], index+1, good, printStr+" *")

	// fmt.Println(temp, parts[index], temp*padding(parts[index])+parts[index])
	correctP2(final, parts, (temp*padding(parts[index]))+parts[index], index+1, good, printStr+" ||")
}

func correct(final int, parts []int) bool {

	n := len(parts)

	for i := range 1 << (n - 1) {
		current := parts[0]
		for j := range n - 1 {
			if i&(1<<j) > 0 {
				current += parts[j+1]
			} else {
				current *= parts[j+1]
			}
		}
		if current == final {

			return true
		}
	}

	return false
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

func scanLine(line []byte) (int, []int) {

	j := string(line)
	lineSplit := strings.Split(j, " ")
	first := strings.TrimSuffix(lineSplit[0], ":")
	others := append([]string{}, lineSplit[1:]...)
	firstNum, err := strconv.Atoi(first)
	noErr(err)

	othersNum := make([]int, 0)
	for _, num := range others {
		temp, err := strconv.Atoi(num)
		noErr(err)
		othersNum = append(othersNum, temp)
	}

	return firstNum, othersNum
}

func padding(n int) int {
	p := 1
	for p < n {
		p *= 10
	}
	return p
}
