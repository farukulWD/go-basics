package main

import "fmt"

/*---------------------------Go String Data Types--------------------------------
The string data type is used to store a sequence of characters (text). String values must be surrounded by double quotes:
-----------------------------------------------------------------------------------------*/

func string_data_type() {
	var s1 string = "Hello"
	var s2 = "Hello"
	var s3 string
	s4 := "Hello"
	fmt.Println("-------------------------------------string-----------------------------------------------------")
	fmt.Println("s1: ", s1) // Returns Hello
	fmt.Println("s2: ", s2) // Returns Hello
	fmt.Println("s3: ", s3) // Returns
	fmt.Println("s4: ", s4) // Returns Hello
}
