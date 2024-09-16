package main

import "fmt"

/*
A program that asks for revenue, expenses and tax rate
calculates earnings before tax, earnings after tax, and ratio (ebt/profit)
Returns EBT, profit and the ratio*/

func main() {

	revenue, expenses, taxRate := getInputFromTerminal()

	earningBeforeTax, earningAfterTax, ratio := calculateOutput(revenue, expenses, taxRate)

	fmt.Printf("Earning before tax: %v\nProfit: %v\nRation: %v", earningBeforeTax, earningAfterTax, ratio)
}

func getInputFromTerminal() (revenue, expenses, taxRate float32) {

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)
	return
}

func calculateOutput(revenue, expenses, taxRate float32) (float32, float32, float32) {
	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - (taxRate / 100))
	ratio := earningBeforeTax / earningAfterTax

	return earningAfterTax, earningBeforeTax, ratio
}
