package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalancedStringSplit(t *testing.T) {
	s := "RLRRLLRLRL"
	expected := 4

	actual := balancedStringSplit(s)
	assert.Equal(t, expected, actual)
}

func balancedStringSplit(s string) int {
	res := 0
	tmp := 0
	for _, v := range s {
		if v == 'R' {
			tmp++
		} else {
			tmp--
		}

		if tmp == 0 {
			res++
		}
	}

	return res
}
