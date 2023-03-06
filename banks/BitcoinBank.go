package banks

import (
	"errors"
	"math/rand"
	"time"
)

type BitcoinBank struct {
	accountBalance float64
}

func NewBitcoinBank(startingAmount float64) *BitcoinBank {
	return &BitcoinBank{
		accountBalance: startingAmount,
	}
}

func (b *BitcoinBank) Balance() float64 {
	return b.accountBalance
}

func (b *BitcoinBank) Deposit(amount float64) {
	b.accountBalance += amount
}

func (b *BitcoinBank) Withdraw(amount float64) error {
	gasFee := generateGasFee(50)
	toRemove := amount + float64(gasFee)
	if b.accountBalance - toRemove < 0 {
		return errors.New("Insufficient funds")
	}
	b.accountBalance -= toRemove
	return nil
}

func generateGasFee(upperbound int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(upperbound)
}
