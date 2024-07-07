package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	nums := []int{2, 11, 10, 1, 3}
	k := 10
	expected := 3

	actual := minOperations(nums, k)

	assert.Equal(t, expected, actual)
}

func minOperations(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			res++
		}
	}

	return res
}