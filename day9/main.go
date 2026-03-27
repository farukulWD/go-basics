package main

// Build a CLI app that stores student names and scores in a Map,
// then prints each student's grade. This practices Maps + loops + if/else.

import "fmt"

func main() {
	scores := map[string]int{
		"Rahim":   90,
		"Sadia":   80,
		"Robiul":  98,
		"Sagor":   70,
		"mr.fail": 40,
	}

	for name, score := range scores {
		fmt.Printf("%s's grade is  %s\n", name, getGrade(score))
	}
}
