package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade string
}

func (s Student) printStudent() {
	fmt.Printf("Name: %s | Age: %d | Grade: %s\n",
		s.Name, s.Age, s.Grade)
}
