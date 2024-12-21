package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	l "github.com/sirupsen/logrus"
)

func main() {
	l.SetLevel(l.DebugLevel)

	file, err := os.Open(os.Args[1])
	noErr(err)
	defer file.Close()

	numbers := readInput(file)

	times := os.Args[2]
	blinks, err := strconv.Atoi(times)
	noErr(err)

	stones := make(map[int]int)

	for _, num := range numbers {
		stones[num] += 1
	}

	fmt.Println("input: ", stones)

	for range blinks {

		current := make(map[int]int)
		for num := range stones {

			c := stones[num]

			if num == 0 {
				current[1] += c
				continue
			}

			if len(fmt.Sprintf("%d", num))%2 == 0 {
				str := fmt.Sprintf("%d", num)
				left := str[:len(str)/2]
				right := str[len(str)/2:]

				leftNum, err := strconv.Atoi(left)
				noErr(err)
				rightNum, err := strconv.Atoi(right)
				noErr(err)

				current[leftNum] += c
				current[rightNum] += c
				continue
			} else {
				current[num*2024] += c
			}

		}
		stones = current
	}

	sum := 0
	for _, v := range stones {
		sum = sum + v
	}
	l.Infoln("number of stones: ", sum)

}

func blink(numbers []int) (out []int) {

	for _, num := range numbers {

		if num == 0 {
			out = append(out, 1)
			continue
		}

		if len(fmt.Sprintf("%d", num))%2 == 0 {
			str := fmt.Sprintf("%d", num)
			first := str[:len(str)/2]
			second := str[len(str)/2:]

			firstNum, err := strconv.Atoi(first)
			noErr(err)
			secNum, err := strconv.Atoi(second)
			noErr(err)

			out = append(out, firstNum, secNum)
			continue
		} else {
			out = append(out, num*2024)
		}

	}

	return
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(r io.Reader) []int {

	var input []int
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		b := scanner.Bytes()
		num, err := strconv.Atoi(string(b))
		noErr(err)
		input = append(input, num)
	}
	return input

}
