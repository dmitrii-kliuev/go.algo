package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCenter(t *testing.T) {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	expected := 2

	actual := findCenter(edges)

	assert.Equal(t, expected, actual)
}

func findCenter(edges [][]int) int {
	a := edges[0][0]
	b := edges[0][1]
	c := edges[1][0]
	d := edges[1][1]

	if a == c || a == d {
		return a
	}

	return b
}

func findCenterV1(edges [][]int) int {
	a := edges[0][0]
	b := edges[0][1]
	aCount := 0
	bCount := 0

	for i := 0; i < len(edges); i++ {
		for j := 0; j < len(edges[0]); j++ {
			if edges[i][j] == a {
				aCount++
				if aCount > 1 {
					return a
				}
			} else if edges[i][j] == b {
				bCount++
				if bCount > 1 {
					return b
				}
			}
		}
	}

	return 0
}
