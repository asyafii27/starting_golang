package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke-", counter)
		counter++
		// counter = counter + 1
		/*
			Perulangan ke- 1
			Perulangan ke- 2
			Perulangan ke- 3
			Perulangan ke- 4
			Perulangan ke- 5
			Perulangan ke- 6
			Perulangan ke- 7
			Perulangan ke- 8
			Perulangan ke- 9
			Perulangan ke- 10
		*/
	}

	// example2
	for counterVersion2 := 1; counterVersion2 <= 10; counterVersion2++ {
		fmt.Println("Perulangan ke-", counterVersion2)

		/*Output
		Perulangan ke- 1
		Perulangan ke- 2
		Perulangan ke- 3
		Perulangan ke- 4
		Perulangan ke- 5
		Perulangan ke- 6
		Perulangan ke- 7
		Perulangan ke- 8
		Perulangan ke- 9
		Perulangan ke- 10
		*/

		/*KETERANGAN
		counterVersion2 := 1; 	=> init statement
		post statement			=> counterVersion2++
		cek statement			=> counterVersion2 <= 10;
		*/
	}

	// FOR RANGE
	slices := []string{"Eko", "Kurniawan", "Khannedy"}
	for i := 0; i < len(slices); i++ {
		fmt.Println(slices[i])
		/*OUTPUT
		Eko
		Kurniawan
		Khannedy*/
	}

	for i, value := range slices {
		fmt.Println("Index", i, "=", value)
		/*Output
		Index 0 = Eko
		Index 1 = Kurniawan
		Index 2 = Khannedy
		*/
	}

	// example2
	// jika tidak menggunakan index
	for _, value := range slices {
		fmt.Println("name =", value)
		/*Output
		name = Eko
		name = Kurniawan
		name = Khannedy
		*/
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "proframmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

	/*Output
	name = Eko
	title = proframmer
	*/

}
