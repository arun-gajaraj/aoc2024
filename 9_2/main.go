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
	// l.Debugln(blockmap)

	for position := len(blockmap) - 1; indexOf(blockmap, ".") <= position; {

		//move file blocks

		block, blockPos := findFileBlockBeforePos(blockmap, position)

		freePos, freeLen := findFirstFreeSpaceOfLen(blockmap, len(block), position+1)

		if freePos == -1 {
			position = blockPos - 1
			continue
		}

		moveFileBlock(blockPos, len(block), freePos, freeLen, &blockmap)

		position = blockPos - 1

	}

	// l.Debugln(blockmap)

	checksum := 0
	for index, b := range blockmap {

		if b == "." {
			continue
		}

		num, err := strconv.Atoi(b)
		noErr(err)

		checksum = checksum + index*num

	}

	l.Infoln("checksum: ", checksum)
}

func moveFileBlock(blockStartPos int, blockLen int, freeStartPos int, _ int, blockmap *[]string) {

	blockStr := (*blockmap)[blockStartPos]

	for i := 0; i < blockLen; i++ {
		(*blockmap)[blockStartPos+i] = "."
	}

	for j := 0; j < blockLen; j++ {
		(*blockmap)[freeStartPos+j] = blockStr
	}
}

func findFileBlockBeforePos(blockmap []string, position int) (block []string, blockPos int) {

	currentFileBlockStarted := false
	currentBlockStr := ""
	currentBlock := make([]string, 0)

	for i := position; i >= 0; i-- {

		if !currentFileBlockStarted && blockmap[i] == "." {
			continue
		}

		if !currentFileBlockStarted && blockmap[i] != "." {
			currentBlockStr = blockmap[i]
			currentFileBlockStarted = true

		}

		if currentFileBlockStarted && blockmap[i] == currentBlockStr {

			currentBlock = append(currentBlock, currentBlockStr)
		}

		if currentFileBlockStarted && blockmap[i] != currentBlockStr {
			blockPos = i + 1
			return currentBlock, blockPos
		}

	}
	return
}

func findFirstFreeSpaceOfLen(blockmap []string, length int, limitPos int) (freePos int, freeLen int) {

	freePos = -1
	currentFreePos := -1
	currentFreePosLen := 0
	freePosStarted := false

	if limitPos > len(blockmap) {
		limitPos = len(blockmap)
	}

	for i := 0; i < limitPos; i++ {

		if !freePosStarted && blockmap[i] == "." {
			freePosStarted = true
			currentFreePos = i
			currentFreePosLen = 1
			continue

		}

		if freePosStarted && blockmap[i] == "." {

			currentFreePosLen++

			if i+1 == limitPos {

				if currentFreePosLen >= length {
					return currentFreePos, currentFreePosLen
				} else {
					return -1, 0
				}
			}

			continue
		}

		if freePosStarted && blockmap[i] != "." {
			if currentFreePosLen >= length {
				return currentFreePos, currentFreePosLen
			} else {

				currentFreePos = -1
				currentFreePosLen = 0
				freePosStarted = false
				continue
			}

		}

	}
	return
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
