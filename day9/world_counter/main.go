package main

import (
	"fmt"
	"strings"
)

// Word Counter

func main() {
	sentence := "go is great go is fast go"
	words := strings.Split(sentence, " ")
	// fmt.Println(words)

	counts := make(map[string]int)

	for _, word := range words {
		// counts[word] = counts[word] + 1
		counts[word] += 1
	}

	for word, count := range counts {
		fmt.Printf("\"%s\" appears %d time(s)\n", word, count)
	}

}
