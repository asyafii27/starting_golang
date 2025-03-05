package main

import "fmt"

func main() {
	var name string
	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	/*Output
	Eko Kurniawan
	Eko Khannedy
	*/

	/*
		Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
		namun, jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak wajib menyebutkan tipe daya variablenya
	*/

	var firstNameSingle = "Eko"
	fmt.Println(firstNameSingle)

	var age = 30
	fmt.Println(age)
	/*Output
	Eko
	30
	*/

	/*
		1. Pda golang, kata kunci var saat membuat variable itu tidak wajib, Asalkan pada saat membuat variable langsung menginisialisasikan datanya
		2. Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan data pada variable tersebut
		3/ hanya untuk deklarasi awal, jika kita ingin mengubah data pada variable tersebut, kita tidak boleh menggunakan kata kunci :=
	*/

	country := "Indonesia"
	fmt.Println(country)

	/* Output
	Indonesia
	*/

	/*
		darasikan multiple variable
	*/

	var (
		firstName = "Eko"
		lastName  = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	/*Output
	Eko
	Khannedy
	*/

}
