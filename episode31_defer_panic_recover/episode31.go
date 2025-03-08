package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	logging()
	fmt.Println("Run application")
}

func main() {
	runApplication()
}
