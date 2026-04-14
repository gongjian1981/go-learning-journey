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

	var students []Student
	for scanner.Scan() {
		line := scanner.Text()
		student, err := ParseStudent(line)
		if err != nil {
			fmt.Printf("Error occurs when parse line : %s\n", line)
			continue
		}
		students = append(students, student)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return students, nil
}

func ParseStudent(line string) (Student, error) {
	var student Student
	_, err := fmt.Sscanf(line, "%s %d %d", &student.Name, &student.Age, &student.Score)
	if err != nil {
		return Student{}, err
	}
	return student, nil
}
