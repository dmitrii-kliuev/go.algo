package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	s := "(1+(2*3)+((8)/4))+1"
	want := 3

	got := maxDepth(s)
	assert.Equal(t, want, got)
}

func maxDepth(s string) int {
	openP := 0
	max := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			openP++
		} else if s[i] == ')' {
			openP--
		}

		if openP > max {
			max = openP
		}
	}

	return max
}
