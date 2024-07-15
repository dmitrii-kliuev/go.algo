package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArithmeticTriplets(t *testing.T) {
	nums := []int{0, 1, 4, 6, 7, 10}
	diff := 3
	want := 2

	got := arithmeticTriplets(nums, diff)
	assert.Equal(t, want, got)
}

func arithmeticTriplets(nums []int, diff int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					res++
				}
			}
		}
	}

	return res
}
