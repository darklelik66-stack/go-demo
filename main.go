package main

import (
	"fmt"
)

func main() {

	const kursUSD = 0.8503
	const kursEUR = 1.18
	const kursRUB_USD = 98.98

	var amountUSD float64 = 100
	var amountEUR float64 = 100
	var resultUSD float64 = amountEUR * kursUSD
	var resultEUR float64 = amountUSD * kursEUR
	var resultRUBEUR float64 = amountEUR * (kursRUB_USD / kursEUR)

	fmt.Println(resultUSD)
	fmt.Println(resultEUR)
	fmt.Println(resultRUBEUR)
}
