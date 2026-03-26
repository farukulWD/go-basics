package main

import "fmt"

/*---------------------------Boolean Data Type--------------------------------

A boolean data type is declared with the bool keyword and can only take the values true or false.

The default value of a boolean data type is false.
-----------------------------------------------------------------------------------------*/

func boolean() {
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value
	fmt.Println("-------------------------------------boolean-----------------------------------------------------")
	fmt.Println("b1: ", b1) // Returns true
	fmt.Println("b2: ", b2) // Returns true
	fmt.Println("b3: ", b3) // Returns false
	fmt.Println("b4: ", b4) // Returns true
}
