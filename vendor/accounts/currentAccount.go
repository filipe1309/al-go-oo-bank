package accounts

type CurrentAccount struct {
	Holder     string
	NumAgency  int
	NumAccount int
	Balance    float64
}

func (c *CurrentAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.Balance
	if canWithdraw {
		c.Balance -= withdrawValue
		return "Withdaw success"
	}
	return "No balance"
}

func (c *CurrentAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.Balance += depositValue
		return "Deposit success", c.Balance
	}
	return "Deposit value < 0", c.Balance
}

func (c *CurrentAccount) Transfer(transferValue float64, destinyAccount *CurrentAccount) bool {
	if transferValue < c.Balance && transferValue > 0 {
		c.Balance -= transferValue
		destinyAccount.Deposit(transferValue)
		return true
	}
	return false
}
