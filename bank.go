package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {
		displayOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			checkBalance(accountBalance)
		} else if choice == 2 {
			accountBalance = depositMoney(accountBalance)
		} else if choice == 3 {
			accountBalance = withdrawMoney(accountBalance)
		} else {
			break
		}
	}
	fmt.Println("Thank you for using Go Bank. Goodbye!")
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
