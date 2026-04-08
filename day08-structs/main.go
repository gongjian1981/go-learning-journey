package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s Student) isPassing() bool {
	return s.Score >= 60
}

func main() {
	fmt.Print("How many students? ")
	var number int
	fmt.Scan(&number)

	if number < 1 {
		fmt.Println("Please enter at least one student.")
		return
	}

	students := make([]Student, 0, number)

	for i := 0; i < number; i++ {
		var name string
		var age int
		var score int

		fmt.Print("Name: ")
		fmt.Scan(&name)

		fmt.Print("Age: ")
		fmt.Scan(&age)

		fmt.Print("Score: ")
		fmt.Scan(&score)

		student := Student{
			Name:  name,
			Age:   age,
			Score: score,
		}
		students = append(students, student)
	}

	printStudents("Student List:", students)

	fmt.Printf("Average Score: %.2f\n\n", averageScore(students))

	top := topStudent(students)
	fmt.Printf("Top Student: %s (Age: %d, Score: %d)\n\n", top.Name, top.Age, top.Score)

	passingStudentsList := passingStudents(students)
	printStudents("Passing Student List:", passingStudentsList)

	sort(students, 0, len(students)-1)
	printStudents("Sorted Student List:", students)
}

func printStudents(label string, students []Student) {
	fmt.Println(label)
	for _, s := range students {
		fmt.Printf("Name: %s\tAge: %d\tScore: %d\n", s.Name, s.Age, s.Score)
	}
	fmt.Println()
}

func averageScore(students []Student) float64 {
	sumScore := 0
	for _, student := range students {
		sumScore += student.Score
	}
	return float64(sumScore) / float64(len(students))
}

func topStudent(students []Student) Student {
	top := students[0]
	for _, student := range students[1:] {
		if student.Score > top.Score {
			top = student
		}
	}
	return top
}

func passingStudents(students []Student) []Student {
	result := make([]Student, 0)
	for _, student := range students {
		if student.isPassing() {
			result = append(result, student)
		}
	}
	return result
}

func sort(students []Student, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(students, left, right)
	sort(students, left, pivot-1)
	sort(students, pivot+1, right)
}

func partition(students []Student, left, right int) int {
	pivot := students[right].Score
	i := left - 1

	for j := left; j < right; j++ {
		if students[j].Score > pivot {
			i++
			students[i], students[j] = students[j], students[i]
		}
	}

	students[i+1], students[right] = students[right], students[i+1]
	return i + 1
}
