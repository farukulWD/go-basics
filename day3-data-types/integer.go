package main

import "fmt"

/*---------------------------Go Integer Data Types--------------------------------
Integer data types are used to store a whole number without decimals, like 35, -50, or 1345000.

The integer data type has two categories:

Signed integers: can store both positive and negative values
Unsigned integers: can only store non-negative values

Tip: The default type for integer is int. If you do not specify a type, the type will be int.
-----------------------------------------------------------------------------------------*/

func integer() {
	var i1 int = 10
	var i2 = 20
	var i3 int
	i4 := 30
	fmt.Println("-------------------------------------integer-----------------------------------------------------")
	fmt.Println("i1: ", i1) // Returns 10
	fmt.Println("i2: ", i2) // Returns 20
	fmt.Println("i3: ", i3) // Returns 0
	fmt.Println("i4: ", i4) // Returns 30
}

// Signed Integers and their range

func signed_integers() {
	fmt.Println("-------------------------------------signed integers-----------------------------------------------------")
	// int — platform-dependent (32 or 64 bit)
	var a int = -2147483648
	fmt.Println("int:   ", a)

	// int8 — 8 bits, range: -128 to 127
	var b int8 = -128
	fmt.Println("int8:  ", b)

	// int16 — 16 bits, range: -32768 to 32767
	var c int16 = 32767
	fmt.Println("int16: ", c)

	// int32 — 32 bits, range: -2147483648 to 2147483647
	var d int32 = 2147483647
	fmt.Println("int32: ", d)

	// int64 — 64 bits, range: -9223372036854775808 to 9223372036854775807
	var e int64 = 9223372036854775807
	fmt.Println("int64: ", e)
}

// unsigned integers and their range

func unsigned_integers() {
	fmt.Println("-------------------------------------unsigned integers-----------------------------------------------------")
	// uint — platform-dependent (32 or 64 bit)
	var a uint = 4294967295
	fmt.Println("uint:   ", a)

	// uint8 — 8 bits, range: 0 to 255
	var b uint8 = 255
	fmt.Println("uint8:  ", b)

	// uint16 — 16 bits, range: 0 to 65535
	var c uint16 = 65535
	fmt.Println("uint16: ", c)

	// uint32 — 32 bits, range: 0 to 4294967295
	var d uint32 = 4294967295
	fmt.Println("uint32: ", d)

	// uint64 — 64 bits, range: 0 to 18446744073709551615
	var e uint64 = 18446744073709551615
	fmt.Println("uint64: ", e)
}
