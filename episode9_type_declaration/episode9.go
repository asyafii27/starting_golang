package main

import "fmt"

func main() {
	type noKTP string
	type married bool

	var noKtpEko noKTP = "1234567890"
	var marriedStatus married = true

	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}

/*
1234567890
true*/
