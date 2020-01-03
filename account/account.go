package account

import "math/big"

type Account struct {
	AccountNumber string
	Balance       *big.Int
}

func NewAccount(accountNumber string, initialBalance *big.Int) *Account {
	return &Account{AccountNumber: accountNumber, Balance: initialBalance}
}
