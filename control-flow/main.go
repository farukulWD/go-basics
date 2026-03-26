package main

import "fmt"

func checkAdultOrNot(age int) {
	if age > 18 {
		fmt.Println("You are adult")
	} else {
		fmt.Println("You are not adult")
	}
}

func getGrade(marks int) string {
	if marks >= 90 {
		return "Grade A"
	} else if marks >= 75 {
		return "Grade B"
	} else if marks >= 50 {
		return "Grade C"
	} else {
		return "Fail"
	}
}

func checkHoliday(day string) {
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Tuesday":
		fmt.Println("Second day of the week")
	case "Friday":
		fmt.Println("Almost weekend!")
	default:
		fmt.Println("Some other day")
	}
}

func forLoop() {
	fmt.Println("Classic for loop:")
	for i := 0; i <= 5; i++ {
		fmt.Print(i, " ")
	}
}

func whileLikeLoop() {
	fmt.Println("While-like loop:")
	i := 0

	for i <= 5 {
		fmt.Print(i, " ")
		i++
	}
}

func main() {
	fmt.Println("===== Control Flow Practice =====")

	// ===============================
	// 1️⃣ if / else
	// ===============================
	checkAdultOrNot(20)

	// ===============================
	// 2️⃣ else if
	// ===============================
	fmt.Println(getGrade(78))
	// ===============================
	// 3️⃣ switch
	// ===============================
	checkHoliday("Friday")

	// ===============================
	// 4️⃣ for loop
	// ===============================
	forLoop()

	fmt.Println()

	whileLikeLoop()
	fmt.Println()

	// ===============================
	// 5️⃣ break & continue
	// ===============================
	fmt.Println("Loop with break and continue:")
	for j := 1; j <= 5; j++ {
		if j == 3 {
			continue // skip 3
		}
		if j == 5 {
			break // stop at 5
		}
		fmt.Print(j, " ")
	}
	fmt.Println()
}
