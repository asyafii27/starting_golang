package main

import "fmt"

func getHello(name string) string {
	word := name + " Hello World"

	return word
}
func main() {
	result := getHello("Assalamualikum")

	fmt.Println(result) //Assalamualikum Hello World
}
