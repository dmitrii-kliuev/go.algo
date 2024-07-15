package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedStream(t *testing.T) {
	os := Constructor(5)

	assert.Equal(t, []string{}, os.Insert(3, "ccccc"))
	assert.Equal(t, []string{"aaaaa"}, os.Insert(1, "aaaaa"))
	assert.Equal(t, []string{"bbbbb", "ccccc"}, os.Insert(2, "bbbbb"))
	assert.Equal(t, []string{}, os.Insert(5, "eeeee"))
	assert.Equal(t, []string{"ddddd", "eeeee"}, os.Insert(4, "ddddd"))
}

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		stream: make([]string, n),
		ptr: 0,
	}
}

func (os *OrderedStream) Insert(idKey int, value string) []string {
	res := make([]string, 0)
	os.stream[idKey-1] = value

	for os.ptr < len(os.stream) && os.stream[os.ptr] != "" {
		res = append(res, os.stream[os.ptr])
		os.ptr++
	}

	return res
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
