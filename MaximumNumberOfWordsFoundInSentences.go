package main

import (
	"fmt"
	"strings"
)

func main() {
	expected := 6
	sentences := []string{
		"alice and bob love leetcode",
		"i think so too",
		"this is great thanks very much",
	}

	actual := mostWordsFound(sentences)

	if actual == expected {
		fmt.Println("Ok")
	} else {
		fmt.Println("Fail")
	}
}

func mostWordsFound(sentences []string) int {
	max := 0
	for _, s := range sentences {
		wordsLen := len(strings.Split(s, " "))

		if wordsLen > max {
			max = wordsLen
		}
	}

	return max
}
