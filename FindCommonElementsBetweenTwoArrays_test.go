package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIntersectionValues(t *testing.T) {
	var tests = []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{2, 3, 2}, []int{1, 2}, []int{2, 1}},
		{[]int{4, 3, 2, 3, 1}, []int{2, 2, 5, 2, 3, 6}, []int{3, 4}},
	}

	for _, test := range tests {
		got := findIntersectionValues(test.nums1, test.nums2)
		assert.Equal(t, test.want, got)
	}
}

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	m1 := map[int]int{}
	m2 := map[int]int{}

	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]]++
	}

	c1 := 0
	for k := range m2 {
		if v, ok := m1[k]; ok {
			c1 += v
		}
	}

	c2 := 0
	for k := range m1 {
		if v, ok := m2[k]; ok {
			c2 += v
		}
	}

	return []int{c1, c2}
}
