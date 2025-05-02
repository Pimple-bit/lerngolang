package main

import "fmt"

func isSortedLexicographically(words []string) bool {
	for i := 1; i < len(words); i++ {
		if words[i-1] > words[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isSortedLexicographically([]string{"apple", "banana", "cherry"})) // true
	fmt.Println(isSortedLexicographically([]string{"apple", "cherry", "banana"})) // false
	fmt.Println(isSortedLexicographically([]string{"a", "aa", "aaa"}))            // true
	fmt.Println(isSortedLexicographically([]string{"zoo", "animal"}))             // false
}
