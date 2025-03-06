package main

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"Novemver",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)      //[Mei Juni Juli]
	fmt.Println(len(slice1)) //3
	fmt.Println(cap(slice1)) //8

	// ini ketika velue nya diubah
	months[5] = "Diubah"
	fmt.Println(slice1) //[Mei Diubah Juli]

	var slice2 = months[10:]
	fmt.Println(slice2) //[Novemver Desember

	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3) //[Novemver Desember Eko]
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3) //[Novemver Bukan Desember Eko]

	fmt.Println(slice2) //[Novemver Desember]
	fmt.Println(months) //[januari Februari Maret April Mei Diubah Juli Agustus September Oktober Novemver Desember]

	//MAKE
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"
	fmt.Println(newSlice) //[Eko Kurniawan]

	// COPY
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice) //[Eko Kurniawan]
}
