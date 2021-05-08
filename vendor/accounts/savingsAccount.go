package accounts

import "customers"

type SavingsAccount struct {
	Holder                           customers.Holder
	NumAgency, NumAccount, Operation int
	balance                          float64
}

func (c *SavingsAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance
	if canWithdraw {
		c.balance -= withdrawValue
		return "Withdaw success"
	}
	return "No balance"
}

func (c *SavingsAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit success", c.balance
	}
	return "Deposit value < 0", c.balance
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
