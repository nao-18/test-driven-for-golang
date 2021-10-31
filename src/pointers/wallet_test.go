package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitconin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitconin(10))
		assertBalance(t, wallet, Bitconin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitconin(20)}
		wallet.Withdraw(Bitconin(10))
		assertBalance(t, wallet, Bitconin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitconin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitconin(100))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}
