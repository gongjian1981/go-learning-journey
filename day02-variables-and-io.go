package main

import "fmt"

func main() {
	var name string
	var age int
	var height float64
	const cmToMeter = 100.0

	fmt.Print("Name : ")
	fmt.Scan(&name)
	fmt.Print("Age : ")
	fmt.Scan(&age)
	fmt.Print("Height (cm): ")
	fmt.Scan(&height)

	heightMeters := height / cmToMeter

	nextYearAge := age + 1

	var i1 int
	var i2 int

	fmt.Print("first integer : ")
	fmt.Scan(&i1)
	fmt.Print("second integer : ")
	fmt.Scan(&i2)

	fmt.Printf("----- Profile -----\nName: %s\nAge: %d (Next year: %d)\nHeight: %.2f m\n\n", name, age, nextYearAge, heightMeters)
	fmt.Printf("----- Calculator -----\na = %d, b = %d\nSum: %d\nDiff: %d\nProduct：%d\n", i1, i2, i1+i2, i1-i2, i1*i2)

	var weight float64
	fmt.Print("Weight (kg): ")
	fmt.Scan(&weight)
	bmi := weight / heightMeters / heightMeters

	fmt.Printf("Your bmi is %.2f\n", bmi)

}
