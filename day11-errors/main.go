package main

import (
	"fmt"
	"math"
)

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}

func main() {

	fmt.Print("Enter two integers: ")
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	result, err := safeDivide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%d / %d = %d\n", a, b, result)
	}

	fmt.Print("Enter a number to find its square root: ")
	var x float64
	_, err = fmt.Scan(&x)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	sqrtResult, err := safeSqrt(x)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Square root of %.2f is %.2f\n", x, sqrtResult)
	}
}

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, MyError{Message: "cannot divide by zero"}
	}
	return a / b, nil
}

func safeSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, MyError{Message: "cannot take square root of a negative number"}
	}
	return math.Sqrt(x), nil
}
