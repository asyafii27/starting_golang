package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"

	return
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)  ///Eko
	fmt.Println(middleName) //Kurniawan
	fmt.Println(lastName)   //H=Khannedy
}
