package main

import (
	"fmt"
	"strings"
	
)

func isUnique(s string) bool {
	s = strings.ToLower(s)

	seen := make(map[rune]bool)

	for _, r := range s {
		if seen[r] {
			return false
		}
		seen[r] = true
	}

	return true
}

func main() {
	tests := []string{
		"abcdef",
		"abcdea",
		"Hello",
		"Привет",
		"абырвалг",
		"абырваАг",
	}

	for _, t := range tests {
		fmt.Printf("%q -> %v\n", t, isUnique(t))
	}
}