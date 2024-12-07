package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeWithDampener(t *testing.T) {
	input := []int{1, 2, 6, 4, 5}

	isSafe := SafeAfterDampening(input)
	assert.True(t, isSafe)

}

func SafeAfterDampening(input []int) bool {

	for i := range input {
		var temp []int

		for j := range input {
			if i == j {
				continue
			}

			temp = append(temp, input[j])
		}

		fmt.Println(temp)
		// Check if safe
		if delta1to3(temp) {
			if isIncreasing(temp) || isDecreasing(temp) {
				return true
			}
		}
	}

	return false
}
