package main

import "fmt"

func main() {
	const usdToEUR = 0.8553
	const usdToRUB = 80.43
	const eurToRUB = usdToRUB/usdToEUR

	fmt.Printf("%.2f RUB\n", eurToRUB)
}