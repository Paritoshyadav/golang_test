package bank

import (
	"reflect"
	"testing"
)

func TestBalance(t *testing.T) {

	t.Run("depositing btc in wallet", func(t *testing.T) {
		wallet := Wallet{}

		amount := Btc(10)

		err := wallet.deposite(amount)

		got := wallet.getBalance()
		want := Btc(10)

		assertwallet(t, err, got, want)

	})

	t.Run("Withdrawing btc without any error wallet", func(t *testing.T) {
		wallet := Wallet{}

		amount := Btc(30)

		wallet.deposite(amount)

		withdraw_amount := Btc(10)

		err := wallet.withdraw(withdraw_amount)

		got := wallet.getBalance()
		want := Btc(20)

		assertwallet(t, err, got, want)

	})
	t.Run("Withdrawing btc with error", func(t *testing.T) {
		wallet := Wallet{}

		amount := Btc(30)

		wallet.deposite(amount)

		withdraw_amount := Btc(100)

		err := wallet.withdraw(withdraw_amount)

		got := wallet.getBalance()
		want := Btc(20)

		asserterror(t, err, got, want)

	})

}

func assertwallet(t testing.TB, err error, got Btc, want Btc) {

	t.Helper()

	if err != nil {
		t.Fatalf("got error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}

func asserterror(t testing.TB, err error, got Btc, want Btc) {

	t.Helper()

	if err == nil {
		t.Fatalf("Should have got the error but we got: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}
