package student

func AverageScore(students []Student) float64 {
	sum := 0
	for _, student := range students {
		sum += student.Score
	}
	return float64(sum) / float64(len(students))
}

func TopStudents(students []Student) []Student {
	var topStudents []Student
	maxScore := 0
	for _, student := range students {
		if student.Score > maxScore {
			maxScore = student.Score
			topStudents = []Student{student}
		} else if student.Score == maxScore {
			topStudents = append(topStudents, student)
		}
	}
	return topStudents
}

func GradeDistribution(students []Student) map[string]int {
	gradeDist := make(map[string]int)
	for _, student := range students {
		switch {
		case student.Score >= 90:
			gradeDist["A"]++
		case student.Score >= 80:
			gradeDist["B"]++
		case student.Score >= 70:
			gradeDist["C"]++
		case student.Score >= 60:
			gradeDist["D"]++
		default:
			gradeDist["F"]++
		}
	}
	return gradeDist
}
