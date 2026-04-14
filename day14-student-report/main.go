package main

import (
	"day14-student-report/student"
	"fmt"
)

func main() {
	students, err := student.ReadStudents("input.txt")
	if err != nil {
		panic(err)
	}
	report := GenerateReport(students)
	err = student.WriteReport("output.txt", report)
	if err != nil {
		panic(err)
	}
}

func GenerateReport(students []student.Student) string {
	if len(students) == 0 {
		return "No students found."
	}
	average := student.AverageScore(students)
	topStudents := student.TopStudents(students)
	gradeDist := student.GradeDistribution(students)
	report := "--- Student Report ---\n\n"
	report += fmt.Sprintf("Average Score: %.2f\n\n", average)
	report += "Top Students:\n"
	for _, s := range topStudents {
		report += fmt.Sprintf("- %s (%d)\n", s.Name, s.Score)
	}
	report += "\n"
	report += "Grade Distribution:\n"
	gradeList := []string{"A", "B", "C", "D", "F"}
	for _, grade := range gradeList {
		report += fmt.Sprintf("%s: %d\n", grade, gradeDist[grade])
	}
	return report
}
