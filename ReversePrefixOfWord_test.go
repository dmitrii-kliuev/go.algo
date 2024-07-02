package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePrefixOfWord(t *testing.T) {
	/*
		Input: word = "abcdefd", ch = "d"
		Output: "dcbaefd"
	*/

	word := "abcdefd"
	var ch byte = 'd'
	expected := "dcbaefd"

	actual := reversePrefix(word, ch)
	assert.Equal(t, expected, actual)
}

func reversePrefix(word string, ch byte) string {
	pos := strings.IndexByte(word, ch)
	if pos == -1 {
		return word
	}

	var buf bytes.Buffer
	for i := pos; i >= 0; i-- {
		buf.WriteByte(word[i])
	}

	buf.WriteString(word[pos+1:])

	return buf.String()
}
