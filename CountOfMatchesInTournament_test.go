package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfMatches(t *testing.T) {
	n := 7
	expected := 6

	actual := numberOfMatches(n)
	assert.Equal(t, expected, actual)
}

func numberOfMatches(n int) int {
	res := 0
	teams := n
	matches := 0
	for teams > 1 {
		if teams%2 != 0 {
			matches = (teams - 1) / 2
			teams = (teams-1)/2 + 1
		} else {
			matches = teams / 2
			teams = teams / 2
		}

		res += matches
	}

	return res
}
