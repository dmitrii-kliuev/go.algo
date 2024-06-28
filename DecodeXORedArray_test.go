package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeXORedArray(t *testing.T) {
	// Input: encoded = [6,2,7,3], first = 4
	// 4
	// 4 ^ 6 = 2
	// 2 ^ 2 = 0
	// 0 ^ 7 = 7
	// 7 ^ 3 = 4
	// Output: [4,2,0,7,4]

	encoded := []int{1, 2, 3}
	first := 1
	expected := []int{1, 0, 2, 1}

	actual := decode(encoded, first)

	assert.Equal(t, expected, actual)
}

func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first

	for i := 0; i < len(encoded); i++ {
		res[i+1] = res[i] ^ encoded[i]
	}

	return res
}
