package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfSteps(t *testing.T) {
	num := 14
	expected := 6

	actual := numberOfSteps(num)

	assert.Equal(t, expected, actual)
}

func numberOfSteps(num int) int {
	s := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		s++
	}

	return s
}
