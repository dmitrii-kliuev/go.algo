package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortPeople(t *testing.T) {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}

	want := []string{"Mary", "Emma", "John"}

	got := sortPeople(names, heights)

	assert.Equal(t, want, got)
}

type Person struct {
	Height int
	Name   string
}

func sortPeople(names []string, heights []int) []string {
	res := make([]string, 0, len(heights))
	persons := make([]Person, len(heights))

	for i := 0; i < len(heights); i++ {
		persons[i] = Person{Height: heights[i], Name: names[i]}
	}

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Height > persons[j].Height
	})

	for _, p := range persons {
		res = append(res, p.Name)
	}

	return res
}

func _(names []string, heights []int) []string {
	res := make([]string, 0, len(names))

	for len(heights) > 0 {

		idx := 0
		tallest := heights[idx]
		for i := 0; i < len(heights); i++ {
			if heights[i] > tallest {
				tallest = heights[i]
				idx = i
			}
		}

		res = append(res, names[idx])

		heights = removeInt(heights, idx)
		names = removeString(names, idx)
	}

	return res
}

func removeInt(nums []int, k int) []int {
	nums[k] = nums[len(nums)-1]
	return nums[:len(nums)-1]
}

func removeString(strs []string, k int) []string {
	strs[k] = strs[len(strs)-1]
	return strs[:len(strs)-1]
}
