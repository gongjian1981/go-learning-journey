package main

import (
	"day13-student-manager/student"
	"fmt"
)

func main() {
	students, err := student.ReadStudents("students.txt")
	if err != nil {
		fmt.Println("Error reading students:", err)
		return
	}

	avg := student.AverageScore(students)
	topStudents := student.TopStudents(students)

	err = student.WriteResult("result.txt", avg, topStudents)
	if err != nil {
		fmt.Println("Error writing result:", err)
		return
	}
}
