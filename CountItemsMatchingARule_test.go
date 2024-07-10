package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountMatches(t *testing.T) {
	ruleKey := "color"
	ruleValue := "silver"

	items := [][]string{
		{"phone", "blue", "pixel"},
		{"computer", "silver", "lenovo"},
		{"phone", "gold", "iphone"},
	}

	expected := 1

	actual := countMatches(items, ruleKey, ruleValue)

	assert.Equal(t, expected, actual)
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	for i := 0; i < len(items); i++ {
		switch ruleKey {
		case "type":
			if items[i][0] == ruleValue {
				res++
			}
		case "color":
			if items[i][1] == ruleValue {
				res++
			}
		case "name":
			if items[i][2] == ruleValue {
				res++
			}
		}
	}

	return res
}
