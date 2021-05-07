package main

import "fmt"

type CurrentAccount struct {
	holder     string
	numAgency  int
	numAccount int
	balance    float64
}

func main() {
	filipesAccount := CurrentAccount{holder: "Filipe", numAgency: 589, numAccount: 123456, balance: 125.50}
	bobsAccount := CurrentAccount{"Bob", 222, 22222, 222.22}
	fmt.Println(filipesAccount, bobsAccount)

	var johnAccount *CurrentAccount
	johnAccount = new(CurrentAccount)
	johnAccount.holder = "John"
	johnAccount.balance = 500
	fmt.Println(*johnAccount)
}
