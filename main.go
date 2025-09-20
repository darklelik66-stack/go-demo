package main

import (
	"fmt"
)

const kursUSD = 0.8503
const kursEUR = 1.18
const kursRUB_USD = 98.98

func main() {

	getUserСurrency()

	var amountUSD float64 = 100
	var amountEUR float64 = 100
	var resultUSD float64 = amountEUR * kursUSD
	var resultEUR float64 = amountUSD * kursEUR
	var resultRUBEUR float64 = amountEUR * (kursRUB_USD / kursEUR)

	fmt.Println(resultUSD)
	fmt.Println(resultEUR)
	fmt.Println(resultRUBEUR)
}

func getUserСurrency() float64 {
	var ammountCurrency float64
	fmt.Println("Введите сумму валюты:")
	fmt.Scan(&ammountCurrency)
	return ammountCurrency
}

func calculateCurrency(ammountCurrency float64, initialCurrency string, finalCurrency string) float64 {

}
