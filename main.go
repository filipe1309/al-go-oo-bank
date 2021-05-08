package main

import (
	"accounts"
	"fmt"
)

func main() {
	johnAccount := accounts.SavingsAccount{}
	johnAccount.Deposit(100)
	johnAccount.Withdraw(555)

	fmt.Println(johnAccount.GetBalance())
}
