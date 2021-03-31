package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		//fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, 10)

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		//fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("Wanted an error but didn't get one")
		}
	})

}
