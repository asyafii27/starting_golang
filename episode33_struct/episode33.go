package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	//cara 1
	var eko Customer
	eko.Name = "Eko Kurniawan"
	eko.Address = "Indonesia"
	eko.Age = 30

	fmt.Println(eko.Name)    //Eko Kurniawan
	fmt.Println(eko.Address) //Inonesia
	fmt.Println(eko.Age)     //30

	//cara 2
	joko := Customer{
		Name:    "Joko",
		Address: "Cirebon",
		Age:     35,
	}

	fmt.Println(joko)

	//cara 3
	budi := Customer{"budi", "Jakarta", 35}
	fmt.Println(budi)

}
