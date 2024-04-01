package pointersanderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	expected := 10

	fmt.Printf("adress of balance in test is %p \n", &wallet.balance)

	if expected != got {
		t.Errorf("expected %d but got %d", expected, got)
	}
}
