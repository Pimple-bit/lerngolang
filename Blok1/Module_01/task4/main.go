package main

import (
	"fmt"
)

func countChars(s string) map[rune]int {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	return charCount
}

func main() {
	str := "съешь ещё этих мягких французских булок, да выпей чаю"
	counts := countChars(str)

	for char, count := range counts {
		fmt.Printf("%c - %d\n", char, count)
	}
}
