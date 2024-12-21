package main

import (
	"bufio"
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
