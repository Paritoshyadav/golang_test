package bank

import (
	"errors"
	"fmt"
)

type Btc float64

type Wallet struct {
	amount Btc
}

//deposite btc method in wallet

func (w *Wallet) deposite(value Btc) error {

	w.amount += value

	return nil

}

//get currrent wallet balance method

func (w *Wallet) getBalance() Btc {

	return w.amount

}

//defining wallet stringer

func (w *Wallet) String() string {
	return fmt.Sprintf("%f Btc", w.amount)
}

//withdraw amount in wallet

func (w *Wallet) withdraw(amount Btc) error {
	if amount > w.getBalance() {
		return errors.New("amount excced the balance")
	}
	w.amount -= amount
	return nil
}
