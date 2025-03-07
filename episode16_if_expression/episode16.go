package main

import "fmt"

func main() {
	var name = "Eko"

	if name == "Eko" {
		fmt.Println("Hello world") //Hello world
	} else if name == "Joko" {
		fmt.Println("Hello world2") //Hello world2
	} else {
		fmt.Println("Hello World3")
	}

	var length = len(name)
	if length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
