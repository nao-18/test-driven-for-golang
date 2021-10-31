package wallet

import "fmt"

type Bitconin int

type Wallet struct {
	balance Bitconin
}

func (b Bitconin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitconin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitconin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitconin) error {
	w.balance -= amount
	return nil
}
