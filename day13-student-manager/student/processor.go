package student

import (
	"bufio"
	"fmt"
	"os"
)

func ReadStudents(filename string) ([]Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	students := make([]Student, 0)

	for scanner.Scan() {
		line := scanner.Text()
		student, err := StringToStudent(line)
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

func WriteResult(filename string, averageScore float64, topStudents []Student) error {
	output := fmt.Sprintf("Average Score: %.2f\n", averageScore)
	for _, student := range topStudents {
		output += fmt.Sprintf("Top Student: %s (%d)\n", student.Name, student.Score)
	}
	return os.WriteFile(filename, []byte(output), 0644)
}
