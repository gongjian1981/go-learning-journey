package main

import "fmt"

func main() {
	fmt.Print("Enter a string: ")
	var str string
	fmt.Scan(&str)

	frequency := charFrequency(str)
	fmt.Println("Frequency:")
	printMap(frequency)

	fmt.Printf("Vowel count: %d\n", countVowels(str))
	fmt.Printf("Reversed: %s\n", reverseString(str))

	fmt.Print("Second string: ")
	var str2 string
	fmt.Scan(&str2)

	fmt.Printf("Anagram: %v\n", isAnagram(str, str2))

	if isAnagramAdvance(str, str2) {
		fmt.Printf("%s and %s are anagrams\n", str, str2)
	} else {
		fmt.Printf("%s and %s are not anagrams\n", str, str2)
	}
}

func printMap(m map[rune]int) {
	for k, v := range m {
		fmt.Printf("%c -> %d\n", k, v)
	}
}

func charFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}
	return freq
}

func countVowels(s string) int {
	count := 0
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
			count++
		}
	}
	return count
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isAnagram(a, b string) bool {
	if len([]rune(a)) != len([]rune(b)) {
		return false
	}

	freq := make(map[rune]int)
	for _, ch := range a {
		freq[ch]++
	}
	for _, ch := range b {
		freq[ch]--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}

func normalizeString(s string) string {
	runes := make([]rune, 0)
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			ch += 'a' - 'A'
		}
		if ch != ' ' {
			runes = append(runes, ch)
		}
	}
	return string(runes)
}

func isAnagramAdvance(a, b string) bool {
	return isAnagram(normalizeString(a), normalizeString(b))
}
