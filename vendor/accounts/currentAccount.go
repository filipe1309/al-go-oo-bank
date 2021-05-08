package accounts

import "customers"

type CurrentAccount struct {
	Holder     customers.Holder
	NumAgency  int
	NumAccount int
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

func (c *CurrentAccount) Transfer(transferValue float64, destinyAccount *CurrentAccount) bool {
	if transferValue < c.balance && transferValue > 0 {
		c.balance -= transferValue
		destinyAccount.Deposit(transferValue)
		return true
	}
	return false
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
