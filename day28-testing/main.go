package main

import "fmt"

func main() {
	fmt.Println("Day 28 - Testing")
	fmt.Println()
	fmt.Println("Run tests:")
	fmt.Println("  go test -v ./day28-testing/service/...")
	fmt.Println("  go test -cover ./day28-testing/service/...")
	fmt.Println("  go test -v -run TestUserService_Login ./day28-testing/service/...")
}
