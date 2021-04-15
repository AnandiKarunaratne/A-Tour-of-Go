package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	wordMap := make(map[string]int)

	for _, key := range words {
		// _, ok := wordMap[key]
		// if ok {
		// 	wordMap[key] += 1
		// } else {
		// 	wordMap[key] = 1
		// }
		wordMap[key] += 1
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
