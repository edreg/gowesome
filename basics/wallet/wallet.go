package wallet

import (
	"errors"
	"fmt"
)

var ErrorInsufficientBalance = errors.New("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// *Pointer - PassByReference
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrorInsufficientBalance
	}
	w.balance -= amount

	return nil
}
