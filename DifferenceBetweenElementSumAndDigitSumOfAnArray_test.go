package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifferenceOfSum(t *testing.T) {
	nums := []int{1, 15, 6, 3}
	expected := 9

	actual := differenceOfSum(nums)

	assert.Equal(t, expected, actual)
}

func differenceOfSum(nums []int) int {
	eSum := 0
	dSum := 0

	for i := 0; i < len(nums); i++ {
		eSum += nums[i]
		dSum += digitSum(nums[i])
	}

	return eSum - dSum
}

func digitSum(num int) int {
	sum := 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}
