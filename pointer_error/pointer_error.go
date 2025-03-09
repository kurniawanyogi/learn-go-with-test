package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientBalance = errors.New("insufficient balance")

func main() {

}

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// kenapa menggunakan pointer *Wallet
// jika tidak menggunakan Pointer maka receiver nya adalah copy dari Wallet
// dalam arti bukan object dari wallet yang terubah ketika Deposit(10)
// when you create a value - like a wallet, it is stored somewhere in memory.
// You can find out what the address of that bit of memory with &myVal.

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	fmt.Printf("address of balance in Deposit() is %p \n", &w.balance)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientBalance
	}

	w.balance -= amount
	return nil
}
