package student

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func StringToStudent(line string) (Student, error) {
	var student Student
	_, err := fmt.Sscanf(line, "%s %d %d", &student.Name, &student.Age, &student.Score)
	if err != nil {
		return Student{}, err
	}
	return student, nil
}

func AverageScore(students []Student) float64 {
	if len(students) == 0 {
		return 0
	}
	total := 0
	for _, student := range students {
		total += student.Score
	}
	return float64(total) / float64(len(students))
}

func TopStudents(students []Student) []Student {
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
