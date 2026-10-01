package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			result := n * n
			fmt.Printf("Квадрат %d = %d\n", n, result)
		}(num)
	}

	wg.Wait()

}
