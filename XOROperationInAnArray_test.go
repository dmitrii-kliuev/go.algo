package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXorOperation(t *testing.T) {
	n := 5
	start := 0
	expected := 8

	actual := xorOperation(n, start)
	assert.Equal(t, expected, actual)
}

func xorOperation(n int, start int) int {
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = start + 2 * i
	}

	res := 0
	for i := 0; i < n; i++ {
		res ^= nums[i]
	}

	return res
}
