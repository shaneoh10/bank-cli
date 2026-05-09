package main

import "fmt"

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
