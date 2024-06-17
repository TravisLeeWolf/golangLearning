package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 5, 12, 35}
	sum := sumUp(1, 24, 12, 4)
	anotherSum := sumUp(numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// Variatic Function
func sumUp(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
