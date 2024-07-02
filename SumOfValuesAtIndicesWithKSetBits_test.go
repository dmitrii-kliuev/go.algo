package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfValuesAtIndicesWithKSetBits(t *testing.T) {
	nums := []int{5, 10, 1, 5, 2}
	k := 1
	expected := 13

	actual := sumIndicesWithKSetBits(nums, k)

	assert.Equal(t, expected, actual)
}

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0

	for i, num := range nums {
		count := 0
		tmp := i

		for j := 0; j < 16; j++ {
			count += tmp & 1
			tmp >>= 1
		}

		if count == k {
			sum += num
		}
	}

	return sum
}
