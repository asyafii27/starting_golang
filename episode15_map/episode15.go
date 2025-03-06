package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)            //map[address:Subang name:Eko title:Programmer]
	fmt.Println(person["name"])    //EKo
	fmt.Println(person["address"]) //Subang

	book := make(map[string]string)
	book["title"] = "belajar Golang"
	book["author"] = "Eko"
	book["ups"] = "Salah"

	fmt.Println(len(book)) //map[author:Eko title:belajar Golang ups:Salah]
	fmt.Println(book)      //3

	delete(book, "ups")

	fmt.Println(len(book)) //2
	fmt.Println(book)      //map[author:Eko title:belajar Golang]
}
