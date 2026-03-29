package main

import "fmt"

/*---------------------Go Structures---------------------*/
/*
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.
*/

type Person struct {
	Name   string
	Age    int
	City   string
	Job    string
	Salary float64
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.Name)
	fmt.Println("Age: ", pers.Age)
	fmt.Println("Job: ", pers.Job)
	fmt.Println("Salary: ", pers.Salary)
}

func (jobHolder Person) fullName() string {
	return jobHolder.Name + " " + jobHolder.Job
}

// product example
type Product struct {
	Name     string
	Price    float64
	Stock    int
	Category string
}

// FinalPrice returns price after discount %
func (p Product) FinalPrice(discountPct float64) float64 {
	return p.Price * (1 - discountPct/100)
}

// StockStatus returns a human-readable stock label
func (p Product) StockStatus() string {
	if p.Stock == 0 {
		return "Out of stock"
	}
	if p.Stock < 5 {
		return "Low stock"
	}
	return "In stock"
}

// Display prints the product listing
func (p Product) Display() {
	fmt.Printf("[%s] %s — ৳%.0f | %s\n",
		p.Category, p.Name, p.FinalPrice(10), p.StockStatus())
}

func main() {
	var person1 Person
	var person2 Person

	person1.Name = "John"
	person1.Age = 30
	person1.City = "New York"
	person1.Job = "Engineer"
	person1.Salary = 100000

	person2.Name = "Jane"
	person2.Age = 25
	person2.City = "London"
	person2.Job = "Doctor"
	person2.Salary = 200000

	// fmt.Println(person1)
	// fmt.Println(person2)

	// call printPerson function
	person3 := Person{
		Name:   "John",
		Age:    30,
		City:   "New York",
		Job:    "Engineer",
		Salary: 100000,
	}

	printPerson(person1)
	printPerson(person2)
	fmt.Println(person1.fullName())
	fmt.Println(person3.fullName())

	// product example
	item := Product{
		Name:     "Laptop",
		Price:    100000,
		Stock:    10,
		Category: "Electronics",
	}

	fmt.Println(item.FinalPrice(10))
	fmt.Println(item.StockStatus())
	item.Display()
}
