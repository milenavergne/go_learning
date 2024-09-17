package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
A program that asks for revenue, expenses and tax rate
calculates earnings before tax, earnings after tax, and ratio (ebt/profit)
Returns EBT, profit and the ratio*/

func main() {

	revenue, err := getFloatFromTerminal("Revenue")
	if err != nil {
		panic(err)
	}
	expenses, err := getFloatFromTerminal("Expenses")
	if err != nil {
		panic(err)
	}
	taxRate, err := getFloatFromTerminal("Tax Rate")
	if err != nil {
		panic(err)
	}

	earningBeforeTax, earningAfterTax, ratio := calculateOutput(revenue, expenses, taxRate)

	fmt.Printf("Earning before tax: %v\nProfit: %v\nRation: %.02f", earningBeforeTax, earningAfterTax, ratio)
	saveValuesInFile("profitCalculator.txt",
		map[string]float32{
			"earningBeforeTax": earningBeforeTax,
			"Profit":           earningAfterTax,
			"ratio":            ratio})
}

func getFloatFromTerminal(text string) (float32, error) {
	var variable float32
	fmt.Print(text + ": ")
	_, err := fmt.Scan(&variable)

	if variable <= 0 {
		return 0, errors.New("inupt must be greater then zero")
	}
	return variable, err
}

func calculateOutput(revenue, expenses, taxRate float32) (float32, float32, float32) {
	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - (taxRate / 100))
	ratio := earningBeforeTax / earningAfterTax

	return earningBeforeTax, earningAfterTax, ratio
}

func saveValuesInFile(fileName string, valueMap map[string]float32) {
	var builder strings.Builder
	for key, value := range valueMap {
		line := fmt.Sprintf("%s: %.2f, ", key, value)
		builder.WriteString(line)
	}
	builder.WriteString("\n")
	os.WriteFile(fileName, []byte(builder.String()), 0644)
}
