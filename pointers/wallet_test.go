package pointers

import (
	"testing"

	"github.com/s4kh/learn-go-tests/util"
)

func assertBalance(t testing.TB, w Wallet, want Eth) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestWall(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Eth(10))

		want := Eth(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Eth(30)}

		err := wallet.Withdraw(Eth(23))
		want := Eth(7)

		if err != nil {
			t.Error("wanted no error but got one")
		}

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startinBalance := Eth(30)
		wallet := Wallet{startinBalance}
		err := wallet.Withdraw(Eth(31))

		assertBalance(t, wallet, startinBalance)
		util.AssertError(t, err, ErrInsufficientFunds)
	})

}
