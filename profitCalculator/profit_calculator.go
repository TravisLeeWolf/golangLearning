package main

import (
	"errors"
	"fmt"
	"os"
)

const resultsFile = "indicators.txt"

// Stores the result of the profit calculator in a txt file
func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile(resultsFile, []byte(results), 0644)
}

// Gets users input data for calculating profit indicators
func getUserInput(prompt string) (float64, error) {
	var value float64
	fmt.Print(prompt + ": ")
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("value cannot be 0 or a negative value")
	}

	return value, nil
}

// Calculates the ebt, profit and ratio
func calculateIndicators(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

// Runs the main program
func main() {
	revenue, err1 := getUserInput("Revenue")
	expenses, err2 := getUserInput("Expenses")
	taxRate, err3 := getUserInput("Tax Rate")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculateIndicators(revenue, expenses, taxRate)

	// fmt.Printf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	fmt.Println("Results stored to file!")
	storeResults(ebt, profit, ratio)
}
