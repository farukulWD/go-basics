package main

func getGrade(score int) string {
	if score >= 90 {
		return "A+"
	} else if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B+"
	} else {
		return "FAIL"
	}
}
