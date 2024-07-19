package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLowerCase(t *testing.T) {
	s := "Hello"
	want := "hello"

	got := toLowerCase(s)
	assert.Equal(t, want, got)
}

func toLowerCase(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			buf.WriteByte(s[i] + 32)
		} else {
			buf.WriteByte(s[i])
		}
	}

	return buf.String()
}
