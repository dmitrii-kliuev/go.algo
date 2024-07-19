package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinalString(t *testing.T) {
	s := "string"
	want := "rtsng"

	got := finalString(s)
	assert.Equal(t, want, got)
}

func finalString(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			buf = *bytes.NewBuffer(reverse(buf))
		} else {
			buf.WriteByte(s[i])
		}
	}

	return buf.String()
}

func reverse(buf bytes.Buffer) []byte {
	p1 := 0
	p2 := buf.Len() - 1
	b := buf.Bytes()
	for i := 0; i < buf.Len()/2; i++ {
		tmp := b[i]
		b[p1] = b[p2]
		b[p2] = tmp
		p1++
		p2--
	}

	return b
}
