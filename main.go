package main

import (
	"fmt"
)

// Currency Map
var rates = map[string]float64{
	"USD": 1.0,
	"INR": 83.12,
	"EUR": 13.0,
	"JPY": 10.0,
}

// Variables Needed
var amount float64
var source_currency string
var target_currency string

// Function for Taking Input
func takeInput() {
	fmt.Println("Enter Amount")
	fmt.Scan(&amount)

	fmt.Println("Enter Source Currency")
	fmt.Scan(&source_currency)

	fmt.Println("Enter Target Currency")
	fmt.Scan(&target_currency)

}

// Function for Main Conversion Logic
func convert() float64 {
	return (amount / rates[source_currency]) * rates[target_currency]
}

// Main Function
func main() {
	fmt.Println("Hello World")

	takeInput()

	var totalSum float64
	totalSum = convert()

	fmt.Println(amount, source_currency, "is equivalent to", totalSum, target_currency)
}
