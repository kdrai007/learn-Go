package main

import (
	"errors"
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(customer *customer, trans transaction) error {

	if trans.transactionType != "deposit" && trans.transactionType != "withdrawal" {
		return errors.New("Unknown! transaction type")
	}

	if trans.transactionType == "deposit" && trans.customerID == customer.id {
		customer.balance += trans.amount
	}

	if trans.transactionType == "withdrawal" && trans.customerID == customer.id {
		if customer.balance < trans.amount {
			return errors.New("insufficient funds")
		}
		customer.balance -= trans.amount
	}
	defer fmt.Println(customer)
	return nil
}

func exerciseTwo() {
	alice := customer{id: 1, balance: 100.0}
	deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

	updateBalance(&alice, deposit)
	// id: 1 balance: 150
}
