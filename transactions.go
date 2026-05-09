package main

import "fmt"

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
