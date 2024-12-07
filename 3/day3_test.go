package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestGetAllMuls(t *testing.T) {

	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)

	}
	defer input.Close()

	content, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}

	expression := `mul\(\d+,\d+\)`

	regex, err := regexp.Compile(expression)
	if err != nil {
		panic(err)
	}

	matches := regex.FindAll(content, -1)

	result := 0

	for _, m := range matches {
		splits := strings.Split(string(m), ",")
		a1 := strings.TrimPrefix(splits[0], "mul(")
		a2 := strings.TrimSuffix(splits[1], ")")

		i1, err := strconv.Atoi(a1)
		if err != nil {
			panic(err)
		}

		i2, err := strconv.Atoi(a2)
		if err != nil {
			panic(err)
		}

		result = result + i1*i2
	}

	fmt.Printf("Result: Uncorrupted numbers multiplied is : %d\n", result)

}

func TestGetAllMulsWithInstruction(t *testing.T) {

	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)

	}
	defer input.Close()

	content, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}

	expression := `mul\(\d+,\d+\)|do\(\)|don't\(\)`

	regex, err := regexp.Compile(expression)
	if err != nil {
		panic(err)
	}

	matches := regex.FindAll(content, -1)

	result := 0
	skipTillDo := false

	for _, m := range matches {

		if string(m) == "don't()" {
			skipTillDo = true
			continue
		} else if string(m) == "do()" {
			skipTillDo = false
			continue
		}

		if skipTillDo {
			continue
		}

		splits := strings.Split(string(m), ",")
		a1 := strings.TrimPrefix(splits[0], "mul(")
		a2 := strings.TrimSuffix(splits[1], ")")

		i1, err := strconv.Atoi(a1)
		if err != nil {
			panic(err)
		}

		i2, err := strconv.Atoi(a2)
		if err != nil {
			panic(err)
		}

		result = result + i1*i2
	}

	fmt.Printf("Result: Uncorrupted numbers multiplied with instruction is : %d\n", result)

}
