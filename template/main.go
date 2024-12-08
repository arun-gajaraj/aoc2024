package main

import (
	"bufio"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	file, err := os.Open("input.txt")
	noErr(err)
	defer file.Close()

	input := readInput(file)
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
