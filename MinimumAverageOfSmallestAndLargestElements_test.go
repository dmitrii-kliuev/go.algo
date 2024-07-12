package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumAverage(t *testing.T) {
	nums := []int{7, 8, 3, 4, 15, 13, 4, 1}
	expected := 5.5

	actual := minimumAverage(nums)

	assert.Equal(t, expected, actual)
}

func minimumAverage(nums []int) float64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	averages := make([]float64, 0)
	p1 := 0
	p2 := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		averages = append(averages, (float64(nums[p1])+float64(nums[p2]))/2)
		p1++
		p2--
	}

	minAverage := averages[0]
	for i := 1; i < len(averages); i++ {
		if averages[i] < minAverage {
			minAverage = averages[i]
		}
	}

	return minAverage
}
