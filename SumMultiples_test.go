package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfMultiples(t *testing.T) {
	n := 7
	expected := 21

	actual := sumOfMultiples(n)

	assert.Equal(t, expected, actual)
}

func sumOfMultiples(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			res += i
		}
	}

	return res
}
