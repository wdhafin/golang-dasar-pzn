package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpDhafin NoKTP = "129312084719"
	var marriedStatus Married = true
	fmt.Println(noKtpDhafin)
	fmt.Println(marriedStatus)
}
