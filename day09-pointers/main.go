package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) addScore(points int) {
	s.Score += points
}

func (s *Student) updateName(newName string) {
	s.Name = newName
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

		students = append(students, Student{
			Name:  name,
			Age:   age,
			Score: score,
		})
	}

	printStudents("Student List:", students)

	fmt.Printf("Before score change: %d\n", students[0].Score)
	changeScoreValue(students[0])
	fmt.Printf("After value function: %d\n", students[0].Score)
	changeScorePointer(&students[0])
	fmt.Printf("After pointer function: %d\n\n", students[0].Score)

	students[0].addScore(5)
	students[0].updateName("NewName")
	printStudents("Updated Student List:", students)

	ptrs := make([]*Student, 0, len(students))
	for i := range students {
		ptrs = append(ptrs, &students[i])
	}

	addBonusToAll(ptrs, 10)
	printStudents("Student List After Add Bonus:", students)
}

func printStudents(label string, students []Student) {
	fmt.Println(label)
	for _, s := range students {
		fmt.Printf("Name: %s\tAge: %d\tScore: %d\n", s.Name, s.Age, s.Score)
	}
	fmt.Println()
}

func addBonusToAll(students []*Student, bonus int) {
	for _, s := range students {
		s.addScore(bonus)
	}
}

func changeScoreValue(s Student) {
	s.Score += 10
}

func changeScorePointer(s *Student) {
	s.Score += 10
}
