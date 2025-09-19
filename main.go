package main

import (
	"fmt"
)

func main() {
	// TODO: use constant for kursUSD and kursEUR
	const kursUSD = 0.8503
	const kursEUR = 1.18

	var amountUSD float64 = 100
	var amountEUR float64 = 100
	var amountUSD float64 = amountEUR * kursUSD
	var amountEUR float64 = amountUSD * kursEUR

	fmt.Println(amountUSD)
	fmt.Println(amountEUR)

}
