package main

import "fmt"

func main() {
	var choice1 = false
	var choice2 = false
	var result = choice1 && choice2

	fmt.Println(result)

	nilaiAkhir := 91
	absensi := 81

	lulusNilaiAkhir := nilaiAkhir > 90
	lulusAbsensi := absensi > 80

	lulus := lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}
