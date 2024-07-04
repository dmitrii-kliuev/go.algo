package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompressRLElist(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{2, 4, 4, 4}

	actual := decompressRLElist(nums)

	assert.Equal(t, expected, actual)
}

func decompressRLElist(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}

	return res
}
