package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	fmt.Printf("Number: %d\n", n)

	var s1 string
	if n == 0 {
		s1 = "ZERO"
	} else if n > 0 {
		s1 = "POSITIVE"
	} else {
		s1 = "NEGATIVE"
	}
	var oddOrEven string

	if n%2 != 0 {
		oddOrEven = "ODD"
	} else {
		oddOrEven = "EVEN"
	}

	fmt.Printf("Type: %s, %s\n", s1, oddOrEven)

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

	fmt.Printf("Your grade is %s.\n", grade)

	if sum := 0; n > 0 {
		three := 0
		fmt.Println("\nSequence:")
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", i)
			sum += i
			if i%3 == 0 {
				three++
			}
		}
		fmt.Printf("\nSum: %d\n", sum)
		fmt.Printf("%d numbers are divisible by 3\n", three)
	}

}
