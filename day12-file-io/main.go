package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func main() {
	students, err := readStudents("students.txt")
	if err != nil {
		fmt.Println("Error reading students:", err)
		return
	}

	avg := averageScore(students)
	topStudents := topStudents(students)

	err = writeResult("result.txt", avg, topStudents)
	if err != nil {
		fmt.Println("Error writing result:", err)
		return
	}
}

func readStudents(filename string) ([]Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	students := make([]Student, 0)

	for scanner.Scan() {
		line := scanner.Text()
		student, err := stringToStudent(line)
		if err != nil {
			fmt.Println("Skipping invalid line:", line)
			continue
		}
		students = append(students, student)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func stringToStudent(line string) (Student, error) {
	var student Student
	_, err := fmt.Sscanf(line, "%s %d %d", &student.Name, &student.Age, &student.Score)
	if err != nil {
		return Student{}, err
	}
	return student, nil
}

func writeResult(filename string, averageScore float64, topStudents []Student) error {
	output := fmt.Sprintf("Average Score: %.2f\n", averageScore)
	for _, student := range topStudents {
		output += fmt.Sprintf("Top Student: %s (%d)\n", student.Name, student.Score)
	}
	return os.WriteFile(filename, []byte(output), 0644)
}

func averageScore(students []Student) float64 {
	if len(students) == 0 {
		return 0
	}
	total := 0
	for _, student := range students {
		total += student.Score
	}
	return float64(total) / float64(len(students))
}

func topStudents(students []Student) []Student {
	if len(students) == 0 {
		return nil
	}

	maxScore := students[0].Score
	for _, student := range students {
		if student.Score > maxScore {
			maxScore = student.Score
		}
	}

	result := make([]Student, 0)
	for _, student := range students {
		if student.Score == maxScore {
			result = append(result, student)
		}
	}
	return result
}
