package main

import "fmt"

type Evaluator interface {
	Evaluate() string
}

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s Student) Evaluate() string {
	switch {
	case s.Score >= 90:
		return "A"
	case s.Score >= 80:
		return "B"
	case s.Score >= 70:
		return "C"
	case s.Score >= 60:
		return "D"
	default:
		return "F"
	}
}

func (s Student) String() string {
	return fmt.Sprintf("Student %s", s.Name)
}

type Course struct {
	Name string
	Avg  int
}

func (c Course) Evaluate() string {
	switch {
	case c.Avg >= 90:
		return "A"
	case c.Avg >= 80:
		return "B"
	case c.Avg >= 70:
		return "C"
	case c.Avg >= 60:
		return "D"
	default:
		return "F"
	}
}

func (c Course) String() string {
	return fmt.Sprintf("Course %s", c.Name)
}

func printEvaluation(e Evaluator) {
	fmt.Printf("%v -> %s\n", e, e.Evaluate())
}

func main() {
	var studentCount, courseCount int

	fmt.Print("How many students? ")
	fmt.Scan(&studentCount)

	fmt.Print("How many courses? ")
	fmt.Scan(&courseCount)

	if studentCount < 0 || courseCount < 0 {
		fmt.Println("Please enter a non-negative number for students and courses.")
		return
	}

	if studentCount+courseCount < 1 {
		fmt.Println("Please enter at least one student or course.")
		return
	}

	students := make([]Student, 0, studentCount)
	for i := 0; i < studentCount; i++ {
		var name string
		var age int
		var score int

		fmt.Print("Student name: ")
		fmt.Scan(&name)

		fmt.Print("Age: ")
		fmt.Scan(&age)

		fmt.Print("Score: ")
		fmt.Scan(&score)

		students = append(students, Student{
			Name:  name,
			Age:   age,
			Score: score,
		})
	}

	courses := make([]Course, 0, courseCount)
	for i := 0; i < courseCount; i++ {
		var name string
		var avg int

		fmt.Print("Course name: ")
		fmt.Scan(&name)

		fmt.Print("Average: ")
		fmt.Scan(&avg)

		courses = append(courses, Course{
			Name: name,
			Avg:  avg,
		})
	}

	evaluators := make([]Evaluator, 0, len(students)+len(courses))
	for _, s := range students {
		evaluators = append(evaluators, s)
	}
	for _, c := range courses {
		evaluators = append(evaluators, c)
	}

	fmt.Println("\nEvaluations:")
	for _, e := range evaluators {
		printEvaluation(e)
	}
}
