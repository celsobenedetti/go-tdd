package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("testing deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("testing withdraw", func(t *testing.T) {
		wallet := Wallet{10.0}

		err := wallet.Withdraw(Bitcoin(6.0))
		if err != nil {
			t.Errorf("wanted no error, but got one")
		}

		assertBalance(t, wallet, Bitcoin(4.0))
	})

	t.Run("testing withdraw insuficient funds", func(t *testing.T) {
		wallet := Wallet{10.0}
		err := wallet.Withdraw(Bitcoin(11.0))

		if err == nil {
			t.Errorf("wanted error when withdraw is bigger then funds")

		}
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
