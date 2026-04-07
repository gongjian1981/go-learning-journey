package main

import "fmt"

func main() {
	fmt.Print("How many numbers?")
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("next number: ")
		fmt.Scan(&nums[i])
	}

	fmt.Print("Numbers: ")
	printArray(nums)

	fmt.Printf("Sum: %d\n", sum(nums))
	fmt.Printf("Max: %d\n", max(nums))
	fmt.Printf("Even count: %d\n", countEven(nums))
	fmt.Print(">10: ")
	printArray(filterGreaterThan10(nums))

	newArray := reverse(nums)
	fmt.Print("Reversed Numbers: ")
	printArray(newArray)
}

func printArray(nums []int) {
	fmt.Print("[")
	for i := range nums {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(nums[i])
	}
	fmt.Println("]")
}

func sum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func max(nums []int) int {
	max := nums[0]
	for i := range nums {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func countEven(nums []int) int {
	evenCount := 0
	for i := range nums {
		if nums[i]%2 == 0 {
			evenCount++
		}
	}
	return evenCount
}

func filterGreaterThan10(nums []int) []int {
	var newArray = make([]int, 0)
	for i := range nums {
		if nums[i] > 10 {
			newArray = append(newArray, nums[i])
		}
	}
	return newArray
}

func reverse(nums []int) []int {
	newArray := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		newArray = append(newArray, nums[i])
	}
	return newArray
}
