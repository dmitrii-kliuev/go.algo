package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotInteger(t *testing.T) {
	n := 8
	want := 6

	got := pivotInteger(n)
	assert.Equal(t, want, got)
}

func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}

	arr := make([]int, 0)
	sum := 0
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
		sum += i
	}

	tmp := 0
	for i := len(arr) - 1; i > 0; i-- {
		tmp += arr[i]
		if tmp == sum {
			return arr[i]
		}
		sum -= arr[i]
	}

	return -1
}
