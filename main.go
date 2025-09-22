package main

import (
	"fmt"
	"strconv"
)

const kursEUR_RUB = 98.88
const kursUSD_RUB = 83.59
const kursEUR_USD = 1.18
const kursUSD_EUR = 0.8493

func main() {
	ammountCurrency, initialCurrency, finalCurrency := getUserСurrency()
	fmt.Println("Получилось " + strconv.FormatFloat(calculateCurrency(ammountCurrency, initialCurrency, finalCurrency), 'f', 1, 64))

}

func getValidCurrency(prompt string) string {
	var currency string
	for {
		fmt.Println(prompt)
		fmt.Scan(&currency)
		if currency == "RUB" || currency == "USD" || currency == "EUR" {
			return currency
		}
		fmt.Println("Неверный ввод. Попробуйте еще раз.")
	}
}

func getValidAmount() float64 {
	var amount float64
	for {
		fmt.Println("Введите сумму валюты:")
		_, err := fmt.Scan(&amount)
		if err == nil && amount > 0 {
			return amount
		}
		fmt.Println("Неверный ввод. Введите положительное число.")
	}
}

func getUserСurrency() (float64, string, string) {
	var initialCurrency, finalCurrency string
	var amountCurrency float64

	for {
		initialCurrency = getValidCurrency("Введите исходную валюту RUB/USD/EUR:")
		finalCurrency = getValidCurrency("Введите валюту для конвертации RUB/USD/EUR:")

		if initialCurrency == finalCurrency {
			fmt.Println("Исходная и конечная валюты не могут быть одинаковыми.")
			fmt.Println("Попробуйте еще раз.")
			continue
		}

		break
	}

	amountCurrency = getValidAmount()

	return amountCurrency, initialCurrency, finalCurrency
}

func calculateCurrency(ammountCurrency float64, initialCurrency string, finalCurrency string) float64 {
	if initialCurrency == "RUB" && finalCurrency == "USD" {
		return ammountCurrency * kursUSD_RUB
	} else if initialCurrency == "RUB" && finalCurrency == "EUR" {
		return ammountCurrency * kursEUR_RUB
	} else if initialCurrency == "USD" && finalCurrency == "RUB" {
		return ammountCurrency * kursUSD_RUB
	} else if initialCurrency == "EUR" && finalCurrency == "RUB" {
		return ammountCurrency * (kursUSD_RUB / kursEUR_RUB)
	} else if initialCurrency == "USD" && finalCurrency == "EUR" {
		return ammountCurrency * kursUSD_EUR
	} else if initialCurrency == "EUR" && finalCurrency == "USD" {
		return ammountCurrency * kursEUR_USD
	}
	return ammountCurrency
}
