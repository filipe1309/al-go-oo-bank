package main

import (
	"accounts"
	"fmt"
)

func main() {
	exAccount := accounts.CurrentAccount{}
	exAccount.Deposit(100)

	fmt.Println(exAccount.GetBalance())
}
