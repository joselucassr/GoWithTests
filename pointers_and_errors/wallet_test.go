package pointersanderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		expected := Bitcoin(10)

		fmt.Printf("adress of balance in test is %p \n", &wallet.balance)

		if expected != got {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		expected := Bitcoin(10)

		if expected != got {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})
}
