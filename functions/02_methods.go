// 02_methods.go demonstrates method declarations and usage in Go, including:
// - Defining methods on custom types
// - Method receivers (pointer vs value)
// - Encapsulation using methods
// - Method call syntax
// - Differences between methods and functions
//
// This file is part of the functions package, which shows different
// ways to work with methods and how they differ from regular functions.
package functions

import "fmt"

// First, we define a data type.
type Wallet struct {
	name    string
	balance int
}

// Now, we attach a function to the 'Wallet' type.
// 'w' is the "receiver". This makes 'Deposit' a METHOD of Wallet.
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// This is also a method of Wallet.
func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Name() string {
	return w.name
}

// DemoMethods runs examples of methods in Go
func DemoMethods() {
	// Create an instance of Wallet
	myWallet := &Wallet{name: "USD", balance: 100}

	// You call the method ON the instance of the wallet.
	fmt.Printf("Methods Example (Wallet - %s):\n", myWallet.Name())
	fmt.Printf("  * Initial balance: $%d\n", myWallet.Balance())

	// Deposit money into the wallet
	myWallet.Deposit(50)
	fmt.Printf("  * After depositing $50, balance: $%d\n", myWallet.Balance())
}
