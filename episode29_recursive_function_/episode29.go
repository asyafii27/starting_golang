package main

import "fmt"

func factoralLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i //result = result * i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		fmt.Println(value)
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factoralLoop(5)
	fmt.Println(loop)              //120
	fmt.Println(5 * 4 * 3 * 2 * 1) //120

	recursive := factorialRecursive(5)
	fmt.Println(recursive) //120

}

// outputnya sama semua
