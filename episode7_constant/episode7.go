package main

import "fmt"

func main() {
	const firstName string = "Eko"
	const lastName = "Khannedy"
	const value = 100

	fmt.Println(firstName)

	// Kalo variable pada constant tidak digunakan, maka tidak akan error

	const (
		country  = "Indonesia"
		city     = "Jakarta"
		district = "Cilandak"
	)
	fmt.Println(country)
}
