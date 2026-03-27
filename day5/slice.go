package main

import "fmt"

// slice_name := []datatype{values}

func slicesExample() {
	fmt.Println(".....................Sliecs example............")
	mySlice := []int{}
	mySlice1 := []int{1, 2, 3}

	fmt.Println("Capacity of the my slice:", cap(mySlice))
	fmt.Println("Length of mySlice array:", len(mySlice1))

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println("Length of myslice2:", len(myslice2))
	fmt.Println("Capacity of myslice2:", cap(myslice2))
	fmt.Println("myslice2:", myslice2)

	// create slice of an array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice3 := arr1[0:4]

	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

}

func makeSlice() {
	fmt.Println(".....................Make Slice example............")
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

	// copy(dst, src)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-13]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Println("numbersCopy:", numbersCopy)
}
