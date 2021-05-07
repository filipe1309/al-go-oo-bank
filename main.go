package main

import "fmt"

type CurrentAccount struct {
	holder     string
	numAgency  int
	numAccount int
	balance    float64
}

func (c *CurrentAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance
	if canWithdraw {
		c.balance -= withdrawValue
		return "Withdaw success"
	}
	return "No balance"
}

func (c *CurrentAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit success", c.balance
	}
	return "Deposit value < 0", c.balance
}

func main() {
	johnAccount := CurrentAccount{}
	johnAccount.holder = "John"
	johnAccount.balance = 500

	fmt.Println(johnAccount.balance)
	status, value := johnAccount.Deposit(2000)
	fmt.Println(status, value)
	fmt.Println(johnAccount.Withdraw(300))
	fmt.Println(johnAccount.balance)
}
