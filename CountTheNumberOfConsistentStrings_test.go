package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountConsistentStrings(t *testing.T) {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	want := 2

	got := countConsistentStrings(allowed, words)
	assert.Equal(t, want, got)
}

func countConsistentStrings(allowed string, words []string) int {
	res := 0
	for _, w := range words {
		if isConsistent(allowed, w) {
			res++
		}
	}

	return res
}

func isConsistent(allowed string, w string) bool {
	for i := 0; i < len(w); i++ {
		if !strings.Contains(allowed, string(w[i])) {
			return false
		}
	}

	return true
}
