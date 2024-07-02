package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSum(t *testing.T) {
	// num := 2932
	// expected := 52

	num := 5592
	expected := 84

	actual := minimumSum(num)

	assert.Equal(t, expected, actual)
}

func minimumSum(num int) int {
	digits := make([]int, 0)

	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}

	sort.Ints(digits)

	return (digits[0] * 10 + digits[2]) + (digits[1] * 10 + digits[3])
}
