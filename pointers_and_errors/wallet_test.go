package pointersanderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	expected := Bitcoin(10)

	fmt.Printf("adress of balance in test is %p \n", &wallet.balance)

	if expected != got {
		t.Errorf("expected %s but got %s", expected, got)
	}
}
