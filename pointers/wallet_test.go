package pointers

import "testing"

func assertBalance(t testing.TB, w Wallet, want Eth) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
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
		assertError(t, err, ErrInsufficientFunds)
	})

}
