package main

import "fmt"

// Interface
// Define the interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 { return 2 * 3.14 * c.Radius }

// Rectangle also implements Shape
type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func printInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

// Interface with a Slice

var shapes = []Shape{
	Circle{Radius: 3},
	Rectangle{Width: 4, Height: 5},
	Circle{Radius: 7},
}

// Stringer Interface — Control How Your Type Prints

type Student struct {
	Name  string
	Score int
}

func (s Student) String() string {
	return fmt.Sprintf("[%s | Score: %d]", s.Name, s.Score)
}

//  Empty Interface — any

func describe(i any) {
	fmt.Printf("Value: %v | Type: %T\n", i, i)
}

// Type switch — handle multiple types
func whatShape(s Shape) {
	switch v := s.(type) {
	case Circle:
		fmt.Println("Circle, radius:", v.Radius)
	case Rectangle:
		fmt.Println("Rectangle:", v.Width, "x", v.Height)
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	fmt.Println("--- Basic Interfaces ---")
	printInfo(Circle{Radius: 5})
	printInfo(Rectangle{Width: 4, Height: 6})

	fmt.Println("\n--- Interface with a Slice ---")
	for i, s := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		whatShape(s)
	}

	fmt.Println("\n--- Stringer Interface (Custom Printer) ---")
	student := Student{"Rahim", 95}
	fmt.Println("Student info:", student.String())

	fmt.Println("\n--- Empty Interface (any) ---")
	describe(42)
	describe("Rahim")

	fmt.Println("\n--- Type Assertion ---")
	var myShape Shape = Circle{Radius: 10}
	concreteCircle, isCircle := myShape.(Circle)
	if isCircle {
		fmt.Println("Radius from concrete type:", concreteCircle.Radius)
	}

	fmt.Println("\n--- Type Switch ---")
	whatShape(myShape)
}
