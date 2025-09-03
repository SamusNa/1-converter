package main

import "fmt"

func main() {
	const usdToEUR = 0.8553
	const usdToRUB = 80.43
	balanceEUR := 100.0
	var balanceRUB float64

	balanceRUB = balanceEUR/usdToEUR*usdToRUB

	fmt.Printf("%.2f RUB\n", balanceRUB)
}