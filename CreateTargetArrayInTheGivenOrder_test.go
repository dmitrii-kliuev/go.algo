package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTargetArray(t *testing.T) {
	/*
		Input: nums = [0,1,2,3,4], index = [0,1,2,2,1]
		Output: [0,4,1,3,2]
	*/
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	expected := []int{0, 4, 1, 3, 2}

	actual := createTargetArray(nums, index)
	assert.Equal(t, expected, actual)
}

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		target = slices.Insert(target, index[i], nums[i])
	}

	return target
}
