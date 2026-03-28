package main

import "fmt"

// incrementByTen adds 10 to the value but since it's passed by value,
// it only affects the local copy within the function.
func incrementByTen(value int) {
	value = value + 10
}

// incrementByTenWithPointer adds 10 to the value using its memory address,
// which modifies the original variable.
func incrementByTenWithPointer(value *int) {
	*value = *value + 10
}

// Student represents a simple student record.
type Student struct {
	Name  string
	Age   int
	Score int
}

// updateStudentScore increases the student's score by 10 points.
func updateStudentScore(student *Student) {
	student.Score += 10
}

func main() {
	// 1. Basic pointer usage
	initialAge := 25
	agePointer := &initialAge
	fmt.Printf("Initial Age: %d, Pointer Value: %d\n", initialAge, *agePointer)

	// 2. Type declaration and value modification via pointer
	currentAge := 21
	var pointerToAge *int = &currentAge
	*pointerToAge = 33
	fmt.Printf("Updated Age (via pointer): %d\n", currentAge)

	// 3. Demonstrating pass by value (no change to original)
	originalValue := 21
	incrementByTen(originalValue)
	fmt.Printf("Value after incrementByTen (should be 21): %d\n", originalValue)

	// 4. Demonstrating pass by reference (changes original)
	updatableValue := 21
	fmt.Printf("Before incrementByTenWithPointer: %d\n", updatableValue)
	incrementByTenWithPointer(&updatableValue)
	fmt.Printf("After incrementByTenWithPointer: %d\n", updatableValue)

	// 5. Pointers with Structs
	studentProfile := Student{Name: "Sara", Age: 21, Score: 86}
	updateStudentScore(&studentProfile)
	fmt.Printf("Updated Student Score: %d\n", studentProfile.Score) // Expected: 96

	// 6. Nil pointer check
	var nilPointer *int
	if nilPointer != nil {
		fmt.Println(*nilPointer)
	} else {
		fmt.Println("Pointer is nil, cannot dereference.")
	}

	// 7. Creating a pointer using new()
	allocatedPointer := new(int)
	fmt.Printf("Value at newly allocated pointer: %d\n", *allocatedPointer)
}
