package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestoreString(t *testing.T) {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	expected := "leetcode"

	actual := restoreString(s, indices)

	assert.Equal(t, expected, actual)
}

func restoreString(s string, indices []int) string {
	a := make([]byte, len(indices))
	for i := 0; i < len(indices); i++ {
		a[indices[i]] = s[i]
	}

	return string(a)
}
