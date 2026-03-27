package main

import "fmt"

// Combines Struct + Map: stores full Student profiles in a map,
// looks one up by name, and prints their details using a method.

func main() {
	db := map[string]Student{
		"sadia": {Name: "Sadia", Age: 45, Grade: "A+"},
		"rahim": {Name: "Rahim", Age: 50, Grade: "A"},
	}

	search := "sadia"
	student, found := db[search]

	if found {
		student.printStudent()
	} else {
		fmt.Println("Student not found!")
	}
}
