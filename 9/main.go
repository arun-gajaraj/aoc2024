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

	input := readInput(file)

	disk := input[0]

	blockmap := make([]string, 0)
	id := 0

	for i := 0; i < len(disk); i++ {
		d := disk[i]
		num, err := strconv.Atoi(string(d))
		noErr(err)

		for range num {
			blockmap = append(blockmap, fmt.Sprintf("%d", id))
		}

		id++
		i++

		if i >= len(disk) {
			break
		}

		d = disk[i]
		num, err = strconv.Atoi(string(d))
		noErr(err)

		for range num {
			blockmap = append(blockmap, ".")
		}

	}
	l.Debugln(blockmap)

	position := len(blockmap) - 1
	for indexOf(blockmap, ".") < position {

		if string(blockmap[position]) != "." {
			numTemp := blockmap[position]
			blockmap[indexOf(blockmap, ".")] = numTemp
			blockmap[position] = "."
		}
		position--

	}

	l.Debugln(blockmap)

	checksum := 0
	for index, b := range blockmap {

		if string(b) == "." {
			break
		}

		num, err := strconv.Atoi(string(b))
		noErr(err)

		checksum = checksum + index*num

	}

	l.Infoln("checksum: ", checksum)
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

	fmt.Println(len(input[0]))
	return input

}

func indexOf(strSlice []string, value string) int {

	for i := range strSlice {
		if strSlice[i] == value {
			return i
		}
	}

	return -1
}
