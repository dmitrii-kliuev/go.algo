package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeMessage(t *testing.T) {
	key := "the quick brown fox jumps ver the lazy dog"
	message := "vkbs bs t suepuv"
	expected := "this is a secret"

	actual := decodeMessage(key, message)
	assert.Equal(t, expected, actual)
}

func decodeMessage(key string, message string) string {
	kMap := make(map[byte]bool)
	uniqueKey := make([]byte, 0)
	for i := 0; i < len(key); i++ {
		if _, ok := kMap[key[i]]; !ok && key[i] != ' ' {
			kMap[key[i]] = true
			uniqueKey = append(uniqueKey, key[i])
		}
	}

	ab := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	res := make([]byte, 0)
	for i := 0; i < len(message); i++ {
		if string(message[i]) == " " {
			res = append(res, ' ')
		} else {
			idx := slices.Index(uniqueKey, message[i])
			res = append(res, byte(ab[idx]))
		}
	}

	return string(res)
}
