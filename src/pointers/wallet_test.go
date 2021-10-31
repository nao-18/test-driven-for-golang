package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitconin(10))
		got := wallet.Balance()
		fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitconin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitconin(20)}
		wallet.Withdraw(Bitconin(10))
		got := wallet.Balance()
		want := Bitconin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
