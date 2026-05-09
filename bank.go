package main

import "fmt"

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = readBalanceFromFile()

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
			writeBalanceToFile(accountBalance)
		case 3:
			accountBalance = withdrawMoney(accountBalance)
			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Thank you for using Go Bank. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
