package arrays

import (
	"testing"

	"github.com/matryer/is"
)

func TestBadBank(t *testing.T) {
	is := is.New(t)

	var (
		riya  = Account{Name: "Riya", Balance: 100.0}
		chris = Account{Name: "Chris", Balance: 75.0}
		adil  = Account{Name: "Adil", Balance: 200.0}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100.0),
			NewTransaction(adil, chris, 25.0),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	is.Equal(newBalanceFor(riya), 200.0)
	is.Equal(newBalanceFor(chris), 0.0)
	is.Equal(newBalanceFor(adil), 175.0)
}
