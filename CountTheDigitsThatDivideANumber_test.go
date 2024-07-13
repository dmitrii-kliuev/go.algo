package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDigits(t *testing.T) {
	num := 7
	expected := 1

	actual := countDigits(num)

	assert.Equal(t, expected, actual)
}

func countDigits(num int) int {
	tmp := num
	res := 0
	for tmp != 0 {
		n := tmp % 10
		if num%n == 0 {
			res++
		}

		tmp /= 10
	}

	return res
}
