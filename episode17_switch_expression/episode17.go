package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, Boleh kenalan")
	}

	// switch case short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch case tanpa kondisi
	lengthV2 := len(name)
	switch {
	case lengthV2 > 10:
		fmt.Println("Nama terlalu panjang")
	case lengthV2 > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
