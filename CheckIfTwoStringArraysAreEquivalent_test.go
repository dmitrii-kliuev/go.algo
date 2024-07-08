package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStringsAreEqual(t *testing.T) {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	expexcted := true

	actual := arrayStringsAreEqual(word1, word2)

	assert.Equal(t, expexcted, actual)
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	w1 := make([]byte, 0)
	for i := 0; i < len(word1); i++ {
		w1 = append(w1, word1[i]...)
	}

	w2 := make([]byte, 0)
	for i := 0; i < len(word2); i++ {
		w2 = append(w2, word2[i]...)
	}

	if len(w1) != len(w2){
		return false
	}

	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			return false
		}
	}

	return true
}
