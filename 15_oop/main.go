package main

import (
	"errors"
	"fmt"
)

var errNegativeBalance = errors.New("ошибка: бфланс не может быть отрицателным")
var errNegativeQuantity = errors.New("ошибка: количество не может быть отрицательным")

type Account struct {
	owner   string
	balance float64
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner}
}

func (a *Account) SetBalance(balance float64) error {
	if balance < 0 {
		return errNegativeBalance
	}

	a.balance = balance
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(quantity float64) error {
	if quantity < 0 {
		return errNegativeQuantity
	}

	a.balance += quantity
	return nil

}

func (a *Account) Withdraw(quantity float64) error {
	if quantity < 0 {
		return errNegativeQuantity
	}

	if a.balance-quantity < 0 {
		return errNegativeBalance
	}

	a.balance -= quantity
	return nil
}

func main() {

	a := Account{}

	fmt.Println(a.GetBalance())

	err := a.Deposit(10)
	if err != nil {
		panic(err)
	}

	err = a.Withdraw(100)
	if err != nil {
		panic(err)
	}

	fmt.Println(a.GetBalance())

}
