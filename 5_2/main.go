package main

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	l "github.com/sirupsen/logrus"
)

type rule struct {
	before int
	after  int
}

type update []int

var rules []rule

func (u update) applies(r rule) bool {
	containsA := false
	containsB := false

	for _, num := range u {
		if num == r.before {
			containsA = true
		}
		if num == r.after {
			containsB = true
		}
	}
	l.Debugf("update %v applies rule %v : %v", u, r, containsA && containsB)
	return containsA && containsB
}

func (u update) obeys(r rule) bool {

	posA, posB := -1, -1

	for i, num := range u {
		if num == r.before {
			posA = i
		}
		if num == r.after {
			posB = i
		}

	}
	l.Debugf("update %v obeys rule %v : %v positions: %d, %d", u, r, posA != -1 && posB != -1 && posA < posB, posA, posB)
	return posA != -1 && posB != -1 && posA < posB
}

func main() {
	l.SetLevel(l.InfoLevel)

	f, err := os.Open("input.txt")
	noErr(err)
	defer f.Close()

	rulesI, updates := getInput(f)
	rules = rulesI

	l.Debugln("rules input: ", rules)
	l.Debugln("updates input: ", updates)

	invalidUpdates := make([]update, 0)

	for _, upd := range updates {
		valid := true
		for _, rul := range rules {

			if upd.applies(rul) && !upd.obeys(rul) {
				valid = false
				break
			}
		}
		if !valid {
			invalidUpdates = append(invalidUpdates, upd)
		}
	}

	// *********** Part 2 *************

	before := new(DefaultSet)

	for _, rul := range rules {
		before.Add(rul.before, rul.after)
	}

	sunOfMiddles := 0
	for _, inv := range invalidUpdates {

		l.Debugln("before sort ", inv)
		sort.Slice(inv, func(i, j int) bool {
			return before.Exists(inv[i], inv[j])
		})

		l.Debugln("after sort ", inv)
		sunOfMiddles = sunOfMiddles + inv[len(inv)/2]
	}

	l.Infoln("solution: ", sunOfMiddles)

}

func getInput(r io.Reader) (rules []rule, updates []update) {
	var inp [][]byte
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lineCopy := make([]byte, len(scanner.Bytes()))
		line := scanner.Bytes()
		copy(lineCopy, line)
		inp = append(inp, lineCopy)
	}
	for _, line := range inp {

		if strings.Contains(string(line), "|") {
			numStrs := strings.Split(string(line), "|")
			before, err := strconv.Atoi(numStrs[0])
			noErr(err)
			after, err := strconv.Atoi(numStrs[1])
			noErr(err)
			rules = append(rules, rule{before, after})

		} else if strings.Contains(string(line), ",") {
			numStrs := strings.Split(string(line), ",")
			var upd update
			for _, numStr := range numStrs {
				num, err := strconv.Atoi(numStr)
				noErr(err)
				upd = append(upd, num)
			}
			updates = append(updates, upd)
		}
	}
	return
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}
