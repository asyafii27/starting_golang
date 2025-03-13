package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHu() {
	fmt.Println("Huu from ", a.Name)
}

func main() {
	//cara 1
	var eko Customer
	eko.Name = "Eko Kurniawan"
	eko.Address = "Indonesia"
	eko.Age = 30

	eko.sayHi("Joko")
	eko.sayHu()

}
