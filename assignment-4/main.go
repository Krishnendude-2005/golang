package main

import (
	"fmt"
)

// Bank Account Structure
type BankAccount struct {
	owner   string
	balance float64
}

func main() {
	fmt.Println("Welcome to Bank Account System")

	//scanning initial values for bank account holder and the balance
	fmt.Println("Enter Initial Details for Bank Account Holder")
	var name string
	var balance float64
	fmt.Println("Enter name : ")
	fmt.Scan(&name)
	fmt.Println("Enter balance: ")
	fmt.Scan(&balance)

	//giving user choices to perfom operations
	fmt.Println("Choose any of the Operations available : \n 1.Display Balance \n 2.Deposit Amount \n 3.Withdraw Amount")

	//any no of operations they want to peform
	fmt.Println("No of operations you want to perform")
	var testcases int
	fmt.Scan(&testcases)

	//creating instance of bankaccount structure
	bankaccount := &BankAccount{
		owner:   name,
		balance: balance,
	}

	for i := range testcases {
		fmt.Println("Operation no ", i+1)
		var operation int
		fmt.Scan(&operation)

		//giving user options to choose
		switch operation {
		case 1:
			bankaccount.DisplayBalance()
		case 2:
			bankaccount.DepositAmount()
		case 3:
			bankaccount.WithdrawAmount()
		}

	}
}

// function for displaying balance
func (b *BankAccount) DisplayBalance() {
	fmt.Println("Owner is ", b.owner, "balance is : ", b.balance)
}

// function for deposit in account
func (b *BankAccount) Deposit(amount float64) {
	b.balance = b.balance + amount
	fmt.Println("Deposit amount is ", amount)
	fmt.Println("Account balance is ", b.balance)
}

// function for withdraw from account
func (b *BankAccount) Withdraw(amount float64) {
	if amount > b.balance {
		fmt.Println("Account balance is too low")
	} else {
		b.balance = b.balance - amount
		fmt.Println("Withdrawl amount is ", amount)
		fmt.Println("Account balance is ", b.balance)
	}
}

// function called in switch case for deposit
func (b *BankAccount) DepositAmount() {
	var amount float64
	fmt.Println("Enter amount to deposit")
	fmt.Scan(&amount)

	b.Deposit(amount)
}

// function called in witch case for withdrawl
func (b *BankAccount) WithdrawAmount() {
	var amount float64
	fmt.Println("Enter amount to withdraw")
	fmt.Scan(&amount)

	b.Withdraw(amount)
}
