package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinBitFlips(t *testing.T) {
	start := 10
	goal := 7

	expected := 3

	actual := minBitFlips(start, goal)

	assert.Equal(t, expected, actual)
}

/*
1010 = 10
0111 = 7
*/
func minBitFlips(start int, goal int) int {
	res := 0
	for i := 0; start != goal; i++ {
		s := start >> i & 1
		g := goal >> i & 1
		if s != g {
			start ^= 1 << i
			res++
		}
	}

	return res
}
