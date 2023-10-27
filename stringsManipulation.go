package main

import (
	"fmt"
	"strings"
)

func testStringsManipulation() {
	// Reverse a string in Go.
	fmt.Printf("[reverseString] hello there: %s\n", reverseString("hello there"))
	fmt.Println("")

	// Implement a function to check if a string is a palindrome.
	fmt.Printf("[isPalindrome] abba, expected true: %v\n", isPalindrome("abba"))
	fmt.Printf("[isPalindrome] abcba, expected true: %v\n", isPalindrome("abcba"))
	fmt.Printf("[isPalindrome] abbta, expected false: %v\n", isPalindrome("abbta"))
	fmt.Printf("[isPalindrome] '', expected true: %v\n", isPalindrome(""))
	fmt.Println("")

	// Recursive
	fmt.Printf("[isPalindromeRecursive] abba, expected true: %v\n", isPalindromeRecursive("abba"))
	fmt.Printf("[isPalindromeRecursive] abcba, expected true: %v\n", isPalindromeRecursive("abcba"))
	fmt.Printf("[isPalindromeRecursive] abbta, expected false: %v\n", isPalindromeRecursive("abbta"))
	fmt.Printf("[isPalindromeRecursive] '', expected true: %v\n", isPalindromeRecursive(""))
	fmt.Println("")

	// Write a function to count the occurrences of a specific character in a string.
	fmt.Printf("[countOccurrences] hello there, e: %d\n", countOccurrences("hello there", "e"))
	fmt.Println("")
}

func reverseString(str string) string {
	splitted := strings.Split(str, "")
	for start, end := 0, len(splitted)-1; end > start; start, end = start+1, end-1 {
		splitted[start], splitted[end] = splitted[end], splitted[start]
	}
	return strings.Join(splitted, "")
}

func isPalindrome(str string) bool {
	splitted := strings.Split(str, "")
	for start, end := 0, len(splitted)-1; end > start; start, end = start+1, end-1 {
		if splitted[start] != splitted[end] {
			return false
		}
	}
	return true
}

func isPalindromeRecursive(str string) bool {
	return isPalindromeRecursiveHelper(strings.Split(str, ""))
}

func isPalindromeRecursiveHelper(str []string) bool {
	if len(str) <= 1 {
		return true
	}

	sLength := len(str) - 1

	return str[0] == str[sLength] && isPalindromeRecursiveHelper(str[1:sLength])
}

func countOccurrences(str, subStr string) int {
	return strings.Count(str, subStr)
}
