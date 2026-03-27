package main

import "fmt"

func myMessage() {
	fmt.Println("I just got executed!")
}

func myFunction(x int, y int) int {
	return x + y
}

// Named return values
func myFunction2(x int, y int) (result int) {
	result = x + y
	return
}

// multiple return values
func myFunction3(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

// recursive functions
func testCount(x int) {
	if x == 11 {
		return
	}
	fmt.Println(x)
	testCount(x + 1)
}

// factorial recursive function
func factorial_recursive(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursive(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	myMessage()
	fmt.Println(myFunction(10, 20))
	fmt.Println("Named return values", myFunction2(10, 20))
	// FIX: Assign multiple values to variables first
	result, txt := myFunction3(10, "Hello")
	fmt.Println("Multiple return values:", result, txt)
	testCount(1)
	fmt.Println("Factorial of 5 is:", factorial_recursive(5))
}
