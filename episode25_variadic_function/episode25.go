package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(1, 1, 1, 1, 2)
	println(total)

	numbers := []int{1, 1, 1, 1, 2}
	total2 := sumAll(numbers...)
	fmt.Println(total2)
}
