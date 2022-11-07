package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Eth float64

type Stringer interface {
	String() string
}

func (e Eth) String() string {
	return fmt.Sprintf("%g ETH", e)
}

type Wallet struct {
	balance Eth
}

func (w *Wallet) Deposit(amount Eth) {
	w.balance += amount
}

func (w *Wallet) Balance() Eth {
	return w.balance
}

func (w *Wallet) Withdraw(amount Eth) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}
