package main

import "fmt"

func main() {
	const usdToEUR = 0.8553
	const usdToRUB = 80.43
	const eurToRUB = usdToRUB/usdToEUR
	baseCurrency := getUserInputBase()
	sum := getUserInputSum()
	toCurrency := getUserInputToCurrency()
	fmt.Printf("%.2f RUB\n", eurToRUB)
}

func converter(baseCurrency string, sum float64, toCurrency string) float64 {

}

func getUserInputBase() string {
	var baseCurrency string
	fmt.Print("Введите исходную валюту (EUR, USD или RUB): ")
	fmt.Scan(&baseCurrency)
	return baseCurrency
}

func getUserInputSum () float64 {
	var sum float64
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&sum)
	return sum
}

func getUserInputToCurrency () string {
	var toCurrency string
	fmt.Print("Введите целевую валюту (EUR, USD или RUB): ")
	fmt.Scan(&toCurrency)
	return toCurrency
}