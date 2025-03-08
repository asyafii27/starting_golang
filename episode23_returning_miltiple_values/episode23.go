package main

import "fmt"

func getFullName() (string, string) {
	return "Eko", "Khannedy"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName) //Eko
	fmt.Println(lastName)  //Khannedy

	//MENGHIRAUKAN RETURN VALUE
	namaDepan, _ := getFullName() //variable yang kedua dihiraukan
	fmt.Println(namaDepan)        //Eko

}
