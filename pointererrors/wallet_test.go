package main

import (
	"errors"
	"fmt"
	"testing"
)

// Bitcoin is a new type based on int, representing Bitcoin amounts.
// Creating new types from existing ones adds domain-specific meaning.
type Bitcoin int

// ErrInsufficientFunds is a package-level variable that represents an error
// for when there are insufficient funds in the wallet.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Wallet is a struct that holds the balance of type Bitcoin.
// It encapsulates the balance, providing methods to interact with it.
type Wallet struct {
	balance Bitcoin
}

// Stringer is an interface that requires the implementation of a String method.
// It allows custom string formatting for types that implement it.
type Stringer interface {
	String() string
}

// String implements the Stringer interface for the Bitcoin type,
// allowing Bitcoin values to be formatted as strings with "BTC".
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit adds the specified amount to the Wallet's balance.
// The method uses a pointer receiver (*Wallet) to modify the original Wallet instance.
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// Withdraw deducts the specified amount from the Wallet's balance.
// If the balance is insufficient, it returns an error.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		// Return an error if there are insufficient funds.
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance returns the current balance of the Wallet.
// This is a simple getter method.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// TestWallet is the main test function. It contains sub-tests for deposit and withdrawal operations.
func TestWallet(t *testing.T) {

	// Sub-test for depositing Bitcoin into the Wallet.
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		// Check if the balance is correct after deposit.
		assertBalance(t, wallet, Bitcoin(10))
	})

	// Sub-test for withdrawing Bitcoin when sufficient funds are available.
	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		// Ensure no error occurred and the balance is correct.
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	// Sub-test for withdrawing Bitcoin when insufficient funds are available.
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		// Ensure the correct error is returned and the balance remains unchanged.
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

// assertBalance checks that the Wallet's balance matches the expected value.
// It uses t.Helper() to mark the function as a helper, improving error reporting.
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

// assertNoError checks that no error occurred when it was not expected.
// If an error is found, the test fails immediately with t.Fatal.
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

// assertError checks that the correct error is returned when one is expected.
// If the error is nil or doesn't match the expected error, the test fails.
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
