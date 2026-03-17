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

	fmt.Print("Your earnings before taxes (EBT) are equal to: ")
	fmt.Println(EBT)
	fmt.Print("Your profit after taxes is equal to: ")
	fmt.Println(profit)
	fmt.Print("Profit to EBT ratio: ")
	fmt.Println(ratio)

}
