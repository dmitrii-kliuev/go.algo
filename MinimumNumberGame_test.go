package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberGame(t *testing.T) {
	nums := []int{5, 4, 2, 3}
	expected := []int{3, 2, 5, 4}

	actual := numberGame(nums)

	assert.Equal(t, expected, actual)
}

func numberGame(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	arr := make([]int, 0)

	for i := 0; i < len(nums); i += 2 {
		arr = append(arr, nums[i+1], nums[i])
	}

	return arr
}
