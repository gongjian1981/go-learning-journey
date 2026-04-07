package main

import "fmt"

func main() {

	fmt.Print("How many numbers?")
	var n int
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Please enter at least one number.")
		return
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("next number: ")
		fmt.Scan(&nums[i])
	}

	fmt.Print("Numbers: ")
	printArray(nums)

	m := buildFrequency(nums)

	fmt.Print("Frequency: \n")
	printMap(m)

	number, max := mostFrequent(m)
	fmt.Printf("Most frequent: %d (count: %d)\n", number, max)

	fmt.Printf("Unique count: %d\n", countUnique(m))

	singles := findSingles(m)
	fmt.Print("Singles: \n")
	printArray(singles)
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

func printMap(m map[int]int) {
	for k, v := range m {
		fmt.Printf("%d -> %d\n", k, v)
	}
}

func buildFrequency(nums []int) map[int]int {
	frequencyCounter := make(map[int]int)
	for i := range nums {
		value, ok := frequencyCounter[nums[i]]
		if ok {
			frequencyCounter[nums[i]] = value + 1
		} else {
			frequencyCounter[nums[i]] = 1
		}
	}
	return frequencyCounter
}

func mostFrequent(freq map[int]int) (int, int) {
	max := 0
	number := 0
	for k, v := range freq {
		if max < v {
			max = v
			number = k
		}
	}
	count := 0
	for _, v := range freq {
		if v == max {
			count++
		}
	}
	if count > 1 {
		fmt.Println("There's more than one number which are most frequent, we only use the number appears the last")
	}
	return number, max
}

func countUnique(freq map[int]int) int {
	return len(freq)
}

func findSingles(freq map[int]int) []int {
	singles := make([]int, 0)
	for k, v := range freq {
		if v == 1 {
			singles = append(singles, k)
		}
	}
	return singles
}
