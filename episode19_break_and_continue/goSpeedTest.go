package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 50000; i++ {
		fmt.Println("Perulangan ke", i)
	}

	duration := time.Since(start).Seconds()  // Hitung durasi dalam detik
	fmt.Println("Waktu eksekusi:", duration) //Waktu eksekusi: 10.7746035
}
