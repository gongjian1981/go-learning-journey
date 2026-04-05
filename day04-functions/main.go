package main

import "fmt"

func main() {

	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	fmt.Printf("Number: %d\n", n)

	sign, parity := getSignAndParity(n)
	fmt.Printf("Type: %s, %s\n", sign, parity)

	fmt.Printf("Your grade is %s.\n", getGrade(n))

	printSequence(n)

	fmt.Printf("\nSum: %d\n", calculateSum(n))
	fmt.Printf("%d numbers are divisible by 3\n", countDivisibleBy3(n))

}

func getSignAndParity(n int) (string, string) {
	var sign string
	if n == 0 {
		sign = "ZERO"
	} else if n > 0 {
		sign = "POSITIVE"
	} else {
		sign = "NEGATIVE"
	}
	var oddOrEven string

	if n%2 != 0 {
		oddOrEven = "ODD"
	} else {
		oddOrEven = "EVEN"
	}
	return sign, oddOrEven
}

func getGrade(n int) string {
	var grade string
	switch {
	case n >= 90:
		grade = "A"
	case n >= 80:
		grade = "B"
	case n >= 70:
		grade = "C"
	default:
		grade = "D"
	}
	return grade
}

func calculateSum(n int) (sum int) {
	sum = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

func countDivisibleBy3(n int) (count int) {
	count = 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			count++
		}
	}
	return
}

func printSequence(n int) {
	for i := 1; i <= n; i++ {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()
}
