package main

import "fmt"

type CurrentAccount struct {
	holder     string
	numAgency  int
	numAccount int
	balance    float64
}

func (c *CurrentAccount) withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance
	if canWithdraw {
		c.balance -= withdrawValue
		return "Success"
	}
	return "No balance"
}

func main() {
	johnAccount := CurrentAccount{}
	johnAccount.holder = "John"
	johnAccount.balance = 500
	fmt.Println(johnAccount.balance)

	fmt.Println(johnAccount.withdraw(300))
	fmt.Println(johnAccount.balance)
}
