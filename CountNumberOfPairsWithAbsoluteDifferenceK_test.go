package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountKDifference(t *testing.T) {
	nums := []int{1, 7, 2, 1, 12}
	k := 5
	want := 2

	got := countKDifference(nums, k)
	assert.Equal(t, want, got)
}

func countKDifference(nums []int, k int) int {
	numsMap := make(map[int]int)
	res := 0

	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] += 1
	}

	for key, count := range numsMap {
		res += count * numsMap[key+k]
	}

	return res
}

func _(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) == float64(k) {
				res++
			}
		}
	}

	return res
}
