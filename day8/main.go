package main

import "fmt"

/*
Maps

A Map stores data as key → value pairs. Think of it like a dictionary — look up a word (key), get its meaning (value).
*/

// Empty map using make()
var scores1 = make(map[string]int)

// Map with values directly
var scores2 = map[string]int{
	"Rahim": 95,
	"Karim": 87,
	"Sadia": 91,
}

func main() {
	fmt.Println(scores1)
	fmt.Println(scores2)

	// iteration maps

	for key, value := range scores2 {
		fmt.Printf("%s : %d\n", key, value)
	}
}
