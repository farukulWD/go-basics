package main

import "fmt"

// ── Declarations with initial values ─────────────────────────────────────────

func variableWithInitialValue() {
	var (
		name      string  = "John"
		age       int     = 30
		height    float64 = 5.9
		isStudent bool    = true
	)
	fmt.Println("With initial value:", name, age, height, isStudent)
}

// ── Declarations without initial values (zero values) ────────────────────────

func variableWithoutInitialValue() {
	var (
		a string
		b int
		c bool
	)
	fmt.Println("Without initial value:", a, b, c)
}

// ── Assign value after declaration ───────────────────────────────────────────

func assignValueAfterDeclaration() {
	var student1 string
	student1 = "John"
	fmt.Println("Assign after declaration:", student1)
}

// ── Multiple variables of the same type ──────────────────────────────────────

func multipleVarsSameType() {
	var a, b, c int = 1, 2, 3
	fmt.Println("Multiple vars, same type:", a, b, c)
}

// ── Multiple variables of different types ────────────────────────────────────

func multipleVarsDifferentTypes() {
	var a, b, c string = "John", "Doe", "Smith"
	fmt.Println("Multiple vars, different types:", a, b, c)
}

// ── Block declaration ─────────────────────────────────────────────────────────

func variableDeclaredInBlock() {
	var (
		name      string  = "John"
		age       int     = 30
		height    float64 = 5.9
		isStudent bool    = true
	)
	fmt.Println("Block declaration:", name, age, height, isStudent)
}

// ── Naming rules ──────────────────────────────────────────────────────────────
//  1. Must start with a letter or underscore (not a digit).
//  2. Can only contain letters, digits, and underscores.
//  3. Case-sensitive.
//  4. Cannot be a reserved keyword.
//
// Common conventions:
//   camelCase  – local variables (preferred in Go)
//   PascalCase – exported (public) identifiers
//   snake_case – avoid in Go; used in some constants by convention elsewhere

var camelCase string = "John"  // unexported
var PascalCase string = "John" // exported
var snake_case string = "John" // non-idiomatic in Go

// ── Entry point ───────────────────────────────────────────────────────────────

func main() {
	variableWithInitialValue()
	variableWithoutInitialValue()
	assignValueAfterDeclaration()
	multipleVarsSameType()
	multipleVarsDifferentTypes()
	variableDeclaredInBlock()

	fmt.Println("camelCase :", camelCase)
	fmt.Println("PascalCase:", PascalCase)
	fmt.Println("snake_case:", snake_case)
}
