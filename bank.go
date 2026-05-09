package main

import (
	"fmt"

	"github.com/shaneoh10/bank-cli/utils"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = utils.ReadFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("No existing balance found. Starting with a balance of 0.00")
		fmt.Println("-------------------------------------")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		displayOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			checkBalance(accountBalance)
		case 2:
			accountBalance = depositMoney(accountBalance)
			utils.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			accountBalance = withdrawMoney(accountBalance)
			utils.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 4:
			fmt.Println("Thank you for using Go Bank. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
