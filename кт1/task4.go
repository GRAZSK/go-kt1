package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println("Исходная:  ", s)
	fmt.Println("Результат: ", reverseWords(s))

	s2 := "день ночь"
	fmt.Println("Исходная:  ", s2)
	fmt.Println("Результат: ", reverseWords(s2))
}