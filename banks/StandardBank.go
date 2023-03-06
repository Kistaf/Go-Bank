package banks

import (
	"errors"
)

type StandardBank struct {
	accountBalance float64
}

func NewStandardBank(startingAmount float64) *StandardBank {
	return &StandardBank{
		accountBalance: startingAmount,
	}
}

func (b *StandardBank) Balance() float64 {
	return b.accountBalance
}

func (b *StandardBank) Deposit(amount float64) {
	b.accountBalance += amount
}

func (b *StandardBank) Withdraw(amount float64) error {
	if b.accountBalance-amount < 0 {
		return errors.New("Insufficient funds")
	}
	b.accountBalance -= amount
	return nil
}
