package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncateSentence(t *testing.T) {
	// s = "Hello how are you Contestant", k = 4
	s := "Hello how are you Contestant"
	k := 4
	expected := "Hello how are you"

	actual :=truncateSentence(s, k)

		assert.Equal(t, expected, actual)
}

func truncateSentence(s string, k int) string {
	splitted := strings.Split(s, " ")
	return strings.Join(splitted[:k], " ")
}
