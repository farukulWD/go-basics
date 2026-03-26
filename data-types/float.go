package main

import "fmt"

/*---------------------------Go Float Data Types--------------------------------
Float data types are used to store a decimal number, like 3.14, -50.5, or 1345000.5.

The float data type has two categories:

Signed floats: can store both positive and negative values
Unsigned floats: can only store non-negative values

Tip: The default type for float is float64. If you do not specify a type, the type will be float64.
-----------------------------------------------------------------------------------------*/

func float() {
	var f1 float64 = 3.14
	var f2 = 3.14
	var f3 float64
	f4 := 3.14
	fmt.Println("-------------------------------------float-----------------------------------------------------")
	fmt.Println("f1: ", f1) // Returns 3.14
	fmt.Println("f2: ", f2) // Returns 3.14
	fmt.Println("f3: ", f3) // Returns 0
	fmt.Println("f4: ", f4) // Returns 3.14
}
