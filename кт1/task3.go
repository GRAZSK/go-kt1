package main

import (
	"fmt"
	
)

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	s := "рыба — дом"
	fmt.Println("Исходная:  ", s)
	fmt.Println("Перевёрнутая:", reverseString(s))


	s2 := "Привет"
	fmt.Println("Исходная:  ", s2)
	fmt.Println("Перевёрнутая:", reverseString(s2))
}