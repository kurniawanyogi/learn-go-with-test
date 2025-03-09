package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertWallet := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	assertError := func(t testing.TB, err, want error) {
		t.Helper()
		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}

		if err != nil {
			if err != want {
				t.Errorf("wanted error %s, got %s", want.Error(), err.Error())
			}
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		want := Bitcoin(10)

		assertWallet(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(100)}
		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(90)

		assertWallet(t, wallet, want)
	})

	t.Run("insufficient balance when withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(110))

		assertError(t, err, ErrInsufficientBalance)
	})
}
