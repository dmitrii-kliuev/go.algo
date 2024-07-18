package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGirstPalindrome(t *testing.T) {
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	want := "ada"

	got := firstPalindrome(words)
	assert.Equal(t, want, got)
}

func firstPalindrome(words []string) string {
	for _, w := range words {
		if IsPalindrome(w) {
			return w
		}
	}

	return ""
}

func IsPalindrome(w string) bool {
	p1 := 0
	p2 := len(w) - 1

	for i := 0; i < len(w)/2; i++ {
		if w[p1] != w[p2] {
			return false
		}
		p1++
		p2--
	}

	return true
}
