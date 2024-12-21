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

	sum := 0

	for _, line := range input {
		result, numbers := scanLine(line)

		good := false
		check(result, numbers, numbers[0], 1, &good)
		if good {
			l.Debugf("%d satisfied by %v", result, numbers)
			sum = sum + result
		}
	}
	fmt.Println("sum: ", sum)
}

// 156: 1 5 6
func check(final int, parts []int, current int, index int, flag *bool) {

	if index == len(parts) {
		if final == current {
			*flag = true
		}
		return
	}

	check(final, parts, current+parts[index], index+1, flag)
	check(final, parts, current*parts[index], index+1, flag)
	check(final, parts, concatNums(current, parts[index]), index+1, flag)
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

func concatNums(a, b int) int {
	firstStr := strconv.Itoa(a)
	secStr := strconv.Itoa(b)

	num, err := strconv.Atoi(firstStr + secStr)
	noErr(err)

	return num
}
