package main

import (
	"accounts"
	"fmt"
)

func PayBoleto(account verifyAccount, boletoValue float64) {
	account.Withdraw(boletoValue)
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {
	johnAccount := accounts.SavingsAccount{}
	johnAccount.Deposit(100)
	PayBoleto(&johnAccount, 60)
	fmt.Println(johnAccount.GetBalance())

	bobAccount := accounts.SavingsAccount{}
	bobAccount.Deposit(500)
	PayBoleto(&bobAccount, 10000)
	fmt.Println(bobAccount.GetBalance())

}
