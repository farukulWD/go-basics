package main

import "fmt"

/*---------------------------Go Array Data Types--------------------------------
Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.
-----------------------------------------------------------------------------------------*/

// Declare an Array

// In Go, there are two ways to declare an array:

// 1. With the `var` keyword:

var withLength = [5]int{1, 2, 3, 4, 5}      // here we define the length of the array
var withoutLength = [...]int{1, 2, 3, 4, 5} // here we infer the length of the array

// add value specefic index

var addValueSpeceficIndex = [5]int{1: 2, 2: 3}

func printArray() {
	fmt.Println("Array with length:", withLength)
	fmt.Println("Array without length:", withoutLength)
	fmt.Println("add value on specefic index", addValueSpeceficIndex)

	// Accessing array elements
	fmt.Println("First element:", withLength[0])
	fmt.Println("Last element:", withLength[4])

	// Modifying array elements
	withLength[0] = 10
	fmt.Println("Modified array:", withLength)

	// Array length
	fmt.Println("Length of array:", len(withLength))

	// Array capacity
	fmt.Println("Capacity of array:", cap(withLength))

	// not initialized
	arr1 := [5]int{}

	fmt.Println("Not initialized:", arr1)

	// iterate the array

	for i := range len(withLength) {
		fmt.Println(i, "Number elements of Array is", withLength[i])
	}

}
