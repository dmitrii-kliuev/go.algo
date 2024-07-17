package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCellsInRange(t *testing.T) {
	s := "K1:M2"
	want := []string{"K1", "K2", "L1", "L2", "M1", "M2"}

	got := cellsInRange(s)
	assert.Equal(t, want, got)
}

/*
"K1" "L1" "M1"
"K2" "L2" "M2"
*/

func cellsInRange(s string) []string {
	colStart := s[0]
	colEnd := s[3]

	rowStart, _ := strconv.Atoi(string(s[1]))
	rowEnd, _ := strconv.Atoi(string(s[4]))

	res := make([]string, 0)

	for i := colStart; i <= colEnd; i++ {
		for j := rowStart; j <= rowEnd; j++ {
			res = append(res, fmt.Sprint(string(i), j))
		}
	}

	return res
}
