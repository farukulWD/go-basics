package main

import "fmt"

/*---------------------------Go hase three basic data types--------------------------------
1. Boolean: represents a boolean value and is either true or false

2. Numeric: represents integer types, floating point values, and complex types

3. String: represents a string value
-----------------------------------------------------------------------------------------*/

func data_types() {
	var is_active bool = true // Boolean
	var age int = 5           // Integer
	var height float32 = 3.14 // Floating point number
	var name string = "Hi!"   // String
	fmt.Println("-------------------------------------data types----------------------------------------------------")
	fmt.Println("Boolean: ", is_active)
	fmt.Println("Integer: ", age)
	fmt.Println("Float:   ", height)
	fmt.Println("String:  ", name)
}
