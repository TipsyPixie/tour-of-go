package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(input string) map[string]int {
	wordCount := make(map[string]int)

	for _, word := range strings.Fields(input) {
		wordCount[word] += 1
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}

