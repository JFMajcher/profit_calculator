package main

import (
	"fmt"
)

func main() {

	var revenue, expenses, taxRate, EBT, profit, ratio float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Extortion rate in [%]: ")
	fmt.Scan(&taxRate)

	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = profit / EBT

	fmt.Println("Your earnings before taxes (EBT) are equal to: ", EBT)
	fmt.Println("Your profit after taxes is equal to: ", profit)
	fmt.Println("Profit to EBT ratio: ", ratio)

	//another way to print values
	fmt.Printf("\nEBT: %v\nProfit: %.2f\nProfit to EBT ratio: %.2f\n\n", EBT, profit, ratio)

	//with formatting before printing:
	formattedEbtMsg := fmt.Sprintf("Earnings before taxes (EBT): %.2f", EBT)
	formattedProfitMsg := fmt.Sprintf("Profit after taxes: %.2f\n", profit)
	formattedRatioMsg := fmt.Sprintf("EBT to profit ratio: %.6f\n", EBT/profit)

	fmt.Println(formattedEbtMsg)
	fmt.Print(formattedProfitMsg, formattedRatioMsg)

	//using `` you can print the strings with the linebreaks and tabulators
	fmt.Printf(`
	Your earnings before taxes (EBT) are equal to: %v
	Profit: %.2f
	Profit to EBT ratio: %.2f\n\n`, EBT, profit, ratio)

}
