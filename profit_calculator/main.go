package main

import "fmt"

/*
A program that asks for revenue, expenses and tax rate
calculates earnings before tax, earnings after tax, and ratio (ebt/profit)
Returns EBT, profit and the ratio*/

func main() {

	revenue := getFloatFromTerminal("Revenue: ")
	expenses := getFloatFromTerminal("Expenses: ")
	taxRate := getFloatFromTerminal("Tax Rate: ")

	earningBeforeTax, earningAfterTax, ratio := calculateOutput(revenue, expenses, taxRate)

	fmt.Printf("Earning before tax: %v\nProfit: %v\nRation: %.02f", earningBeforeTax, earningAfterTax, ratio)
}

func getFloatFromTerminal(text string) float32 {
	var variable float32
	fmt.Print(text)
	fmt.Scan(&variable)
	return variable
}

func calculateOutput(revenue, expenses, taxRate float32) (float32, float32, float32) {
	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - (taxRate / 100))
	ratio := earningBeforeTax / earningAfterTax

	return earningBeforeTax, earningAfterTax, ratio
}
