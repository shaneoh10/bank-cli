package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

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

func displayOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

func checkBalance(balance float64) {
	fmt.Printf("Your balance is %.2f\n", balance)
}

func depositMoney(balance float64) float64 {
	var depositAmount float64

	fmt.Print("Enter amount to deposit: ")
	fmt.Scanln(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Invalid deposit amount.")
		return balance
	}

	balance += depositAmount

	fmt.Printf("You have deposited %.2f. Your new balance is %.2f\n", depositAmount, balance)

	return balance
}

func withdrawMoney(balance float64) float64 {
	var withdrawalAmount float64

	fmt.Print("Enter amount to withdraw: ")
	fmt.Scanln(&withdrawalAmount)

	if withdrawalAmount <= balance {
		balance -= withdrawalAmount
		fmt.Printf("You have withdrawn %.2f. Your new balance is %.2f\n", withdrawalAmount, balance)
	} else {
		fmt.Println("Insufficient funds.")
	}

	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0.0, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0.0, errors.New("Failed to parse balance from file.")
	}

	return balance, nil
}
